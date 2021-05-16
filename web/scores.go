// Copyright 2020 Team 254. All Rights Reserved.
// Author: kenschenke@gmail.com (Ken Schenke)
//
// Web handlers for handling realtime scores API.

/*

API Docs

JSON Schema:

{
   “red”: {“auto”: 99, “teleop”: 99, “endgame": 99},
   “blue”: {“auto”: 99, “teleop”: 99, “endgame": 99}
}

GET http://10.0.100.5/api/scores

Returns current score.

PUT http://10.0.100.5/api/scores

Sets the current scores from the request body. All
parts are optional. Anything missing is set to zero.

Example:

{
   “red”: {“auto”: 10}
}

Red teleop and endgame are set to zero as well as all blue scores.

PATCH http://10.0.100.5/api/scores

Adds or subtracts the current scores from the request
body. All parts are optional. Scores missing from the
request body are left untouched.

Example:

{
   “red”: {“auto”: 10},
   "blue": {"teleop": -5}
}

10 is added to red auto. Red teleop and endgame are left untouched.
5 is subtracted from blue teleop. Blue auto and endgame are left untouched.

*/

package web

import (
	"encoding/json"
	"github.com/Team254/cheesy-arena-lite/field"
	"github.com/Team254/cheesy-arena-lite/game"
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
	if web.arena.MatchState == field.PreMatch || web.arena.MatchState == field.TimeoutActive ||
		web.arena.MatchState == field.PostTimeout {
		http.Error(w, "Score cannot be updated in this match state", http.StatusBadRequest)
		return
	}

	var scores jsonScore
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleWebErr(w, err)
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
}
