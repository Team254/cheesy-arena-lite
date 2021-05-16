// Copyright 2020 Team 254. All Rights Reserved.
// Author: kenschenke@gmail.com (Ken Schenke)

package web

import (
	"encoding/json"
	"github.com/Team254/cheesy-arena-lite/field"
	"github.com/Team254/cheesy-arena-lite/game"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetScores(t *testing.T) {
	web := setupTestWeb(t)

	score1 := game.TestScore1()
	score2 := game.TestScore2()
	web.arena.RedScore.AutoPoints = score1.AutoPoints
	web.arena.RedScore.TeleopPoints = score1.TeleopPoints
	web.arena.RedScore.EndgamePoints = score1.EndgamePoints
	web.arena.BlueScore.AutoPoints = score2.AutoPoints
	web.arena.BlueScore.TeleopPoints = score2.TeleopPoints
	web.arena.BlueScore.EndgamePoints = score2.EndgamePoints

	recorder := web.getHttpResponse("/api/scores")
	assert.Equal(t, 200, recorder.Code)

	var reqScores jsonScore
	json.Unmarshal(recorder.Body.Bytes(), &reqScores)
	assert.Equal(t, score1.AutoPoints, reqScores.Red.Auto)
	assert.Equal(t, score1.TeleopPoints, reqScores.Red.Teleop)
	assert.Equal(t, score1.EndgamePoints, reqScores.Red.Endgame)
	assert.Equal(t, score2.AutoPoints, reqScores.Blue.Auto)
	assert.Equal(t, score2.TeleopPoints, reqScores.Blue.Teleop)
	assert.Equal(t, score2.EndgamePoints, reqScores.Blue.Endgame)
}

func TestPatchScores(t *testing.T) {
	web := setupTestWeb(t)
	var recorder *httptest.ResponseRecorder

	web.arena.MatchState = field.PreMatch
	recorder = web.patchHttpResponse("/api/scores", "{}")
	assert.Equal(t, 400, recorder.Code)
	assert.Equal(t, "Score cannot be updated in this match state\n", recorder.Body.String())

	score1 := game.TestScore1()
	score2 := game.TestScore2()
	web.arena.RedScore.AutoPoints = score1.AutoPoints
	web.arena.RedScore.TeleopPoints = score1.TeleopPoints
	web.arena.RedScore.EndgamePoints = score1.EndgamePoints
	web.arena.BlueScore.AutoPoints = score2.AutoPoints
	web.arena.BlueScore.TeleopPoints = score2.TeleopPoints
	web.arena.BlueScore.EndgamePoints = score2.EndgamePoints

	web.arena.MatchState = field.PostMatch
	recorder = web.patchHttpResponse("/api/scores",
		"{\"red\":{\"auto\":5,\"teleop\":10,\"endgame\":15}}")
	assert.Equal(t, 200, recorder.Code)

	assert.Equal(t, score1.AutoPoints+5, web.arena.RedScore.AutoPoints)
	assert.Equal(t, score1.TeleopPoints+10, web.arena.RedScore.TeleopPoints)
	assert.Equal(t, score1.EndgamePoints+15, web.arena.RedScore.EndgamePoints)
	assert.Equal(t, score2.AutoPoints, web.arena.BlueScore.AutoPoints)
	assert.Equal(t, score2.TeleopPoints, web.arena.BlueScore.TeleopPoints)
	assert.Equal(t, score2.EndgamePoints, web.arena.BlueScore.EndgamePoints)

	recorder = web.patchHttpResponse("/api/scores",
		"{\"blue\":{\"auto\":-5,\"teleop\":-10,\"endgame\":-15}}")
	assert.Equal(t, 200, recorder.Code)

	assert.Equal(t, score1.AutoPoints+5, web.arena.RedScore.AutoPoints)
	assert.Equal(t, score1.TeleopPoints+10, web.arena.RedScore.TeleopPoints)
	assert.Equal(t, score1.EndgamePoints+15, web.arena.RedScore.EndgamePoints)
	assert.Equal(t, score2.AutoPoints-5, web.arena.BlueScore.AutoPoints)
	assert.Equal(t, score2.TeleopPoints-10, web.arena.BlueScore.TeleopPoints)
	assert.Equal(t, score2.EndgamePoints-15, web.arena.BlueScore.EndgamePoints)
}

func TestPutScores(t *testing.T) {
	web := setupTestWeb(t)
	var recorder *httptest.ResponseRecorder

	web.arena.MatchState = field.PreMatch
	recorder = web.putHttpResponse("/api/scores", "{}")
	assert.Equal(t, 400, recorder.Code)
	assert.Equal(t, "Score cannot be updated in this match state\n", recorder.Body.String())

	score1 := game.TestScore1()
	score2 := game.TestScore2()
	web.arena.RedScore.AutoPoints = score1.AutoPoints
	web.arena.RedScore.TeleopPoints = score1.TeleopPoints
	web.arena.RedScore.EndgamePoints = score1.EndgamePoints
	web.arena.BlueScore.AutoPoints = score2.AutoPoints
	web.arena.BlueScore.TeleopPoints = score2.TeleopPoints
	web.arena.BlueScore.EndgamePoints = score2.EndgamePoints

	web.arena.MatchState = field.PostMatch
	recorder = web.putHttpResponse("/api/scores",
		"{\"red\":{\"auto\":5,\"teleop\":10,\"endgame\":15}}")
	assert.Equal(t, 200, recorder.Code)

	assert.Equal(t, 5, web.arena.RedScore.AutoPoints)
	assert.Equal(t, 10, web.arena.RedScore.TeleopPoints)
	assert.Equal(t, 15, web.arena.RedScore.EndgamePoints)
	assert.Equal(t, 0, web.arena.BlueScore.AutoPoints)
	assert.Equal(t, 0, web.arena.BlueScore.TeleopPoints)
	assert.Equal(t, 0, web.arena.BlueScore.EndgamePoints)

	recorder = web.putHttpResponse("/api/scores",
		"{\"blue\":{\"auto\":5,\"teleop\":10,\"endgame\":15}}")
	assert.Equal(t, 200, recorder.Code)

	assert.Equal(t, 0, web.arena.RedScore.AutoPoints)
	assert.Equal(t, 0, web.arena.RedScore.TeleopPoints)
	assert.Equal(t, 0, web.arena.RedScore.EndgamePoints)
	assert.Equal(t, 5, web.arena.BlueScore.AutoPoints)
	assert.Equal(t, 10, web.arena.BlueScore.TeleopPoints)
	assert.Equal(t, 15, web.arena.BlueScore.EndgamePoints)
}
