// Copyright 2020 Team 254. All Rights Reserved.
// Author: kenschenke@gmail.com (Ken Schenke)
//
// Web handlers for handling realtime scores API.

package web

import (
	"encoding/json"
	"fmt"
	"github.com/Team254/cheesy-arena/field"
	"github.com/Team254/cheesy-arena/game"
	"io/ioutil"
	"net/http"
)

type jsonAllianceScore struct {
	Auto    int `json:"auto"`
	Teleop  int `json:"teleop"`
	Endgame int `json:"endgame"`
}

type jsonScore struct {
	Red  jsonAllianceScore `json:"red"`
	Blue jsonAllianceScore `json:"blue"`
}

func (web *Web) getScoresHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(jsonScore{
		Red: jsonAllianceScore{
			Auto:    web.arena.RedScore.AutoPoints,
			Teleop:  web.arena.RedScore.TeleopPoints,
			Endgame: web.arena.RedScore.EndgamePoints,
		},
		Blue: jsonAllianceScore{
			Auto:    web.arena.BlueScore.AutoPoints,
			Teleop:  web.arena.BlueScore.TeleopPoints,
			Endgame: web.arena.BlueScore.EndgamePoints,
		},
	})
}

func (web *Web) setScoresHandler(w http.ResponseWriter, r *http.Request) {
	if web.arena.MatchState == field.PreMatch || web.arena.MatchState == field.TimeoutActive || web.arena.MatchState == field.PostTimeout {
		fmt.Fprintf(w, "Score cannot be updated in this match state")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	var scores jsonScore
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Score data missing")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.Unmarshal(reqBody, &scores)

	if r.Method == "PUT" {
		web.arena.RedScore = new(game.Score)
		web.arena.BlueScore = new(game.Score)
	}

	web.arena.RedScore.AutoPoints += scores.Red.Auto
	web.arena.RedScore.TeleopPoints += scores.Red.Teleop
	web.arena.RedScore.EndgamePoints += scores.Red.Endgame
	web.arena.BlueScore.AutoPoints += scores.Blue.Auto
	web.arena.BlueScore.TeleopPoints += scores.Blue.Teleop
	web.arena.BlueScore.EndgamePoints += scores.Blue.Endgame
	web.arena.RealtimeScoreNotifier.Notify()

	w.WriteHeader(http.StatusOK)
}
