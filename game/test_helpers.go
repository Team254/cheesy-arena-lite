// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Helper methods for use in tests in this package and others.

package game

func TestScore1() *Score {
	return &Score{
		AutoPoints:    45,
		TeleopPoints:  80,
		EndgamePoints: 30,
	}
}

func TestScore2() *Score {
	return &Score{
		AutoPoints:    15,
		TeleopPoints:  40,
		EndgamePoints: 25,
	}
}

func TestRanking1() *Ranking {
	return &Ranking{254, 1, 0, RankingFields{20, 625, 90, 554, 0.254, 3, 2, 1, 10}}
}

func TestRanking2() *Ranking {
	return &Ranking{1114, 2, 1, RankingFields{18, 700, 625, 90, 0.1114, 1, 3, 2, 10}}
}
