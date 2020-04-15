// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreSummary(t *testing.T) {
	redScore := TestScore1()
	blueScore := TestScore2()

	redSummary := redScore.Summarize()
	assert.Equal(t, 45, redSummary.AutoPoints)
	assert.Equal(t, 80, redSummary.TeleopPoints)
	assert.Equal(t, 30, redSummary.EndgamePoints)

	blueSummary := blueScore.Summarize()
	assert.Equal(t, 15, blueSummary.AutoPoints)
	assert.Equal(t, 40, blueSummary.TeleopPoints)
	assert.Equal(t, 25, blueSummary.EndgamePoints)
}

func TestScoreEquals(t *testing.T) {
	score1 := TestScore1()
	score2 := TestScore1()
	assert.True(t, score1.Equals(score2))
	assert.True(t, score2.Equals(score1))

	score3 := TestScore2()
	assert.False(t, score1.Equals(score3))
	assert.False(t, score3.Equals(score1))

	score2 = TestScore1()
	score2.AutoPoints = 20
	assert.False(t, score1.Equals(score2))
	assert.False(t, score2.Equals(score1))

	score2 = TestScore1()
	score2.TeleopPoints = 35
	assert.False(t, score1.Equals(score2))
	assert.False(t, score2.Equals(score1))

	score2 = TestScore1()
	score2.EndgamePoints = 15
	assert.False(t, score1.Equals(score2))
	assert.False(t, score2.Equals(score1))
}
