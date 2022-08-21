// Copyright 2020 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model representing the instantaneous score of a match.

package game

type Score struct {
	AutoPoints    int
	TeleopPoints  int
	EndgamePoints int
}

// Calculates and returns the summary fields used for ranking and display.
func (score *Score) Summarize() *ScoreSummary {
	summary := new(ScoreSummary)

	summary.AutoPoints = score.AutoPoints
	summary.TeleopPoints = score.TeleopPoints
	summary.EndgamePoints = score.EndgamePoints
	summary.Score = summary.AutoPoints + summary.TeleopPoints + summary.EndgamePoints

	return summary
}

// Returns true if and only if all fields of the two scores are equal.
func (score *Score) Equals(other *Score) bool {
	if score.AutoPoints != other.AutoPoints ||
		score.TeleopPoints != other.TeleopPoints ||
		score.EndgamePoints != other.EndgamePoints {
		return false
	}

	return true
}
