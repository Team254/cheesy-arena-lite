// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Functions for calculating the qualification rankings.

package tournament

import (
	"github.com/Team254/cheesy-arena-lite/game"
	"github.com/Team254/cheesy-arena-lite/model"
	"sort"
)

// Determines the rankings from the stored match results, and saves them to the database.
func CalculateRankings(database *model.Database, preservePreviousRank bool) (game.Rankings, error) {
	matches, err := database.GetMatchesByType("qualification")
	if err != nil {
		return nil, err
	}
	rankings := make(map[int]*game.Ranking)
	for _, match := range matches {
		if !match.IsComplete() {
			continue
		}
		matchResult, err := database.GetMatchResultForMatch(match.Id)
		if err != nil {
			return nil, err
		}
		if !match.Red1IsSurrogate {
			addMatchResultToRankings(rankings, match.Red1, matchResult, true)
		}
		if !match.Red2IsSurrogate {
			addMatchResultToRankings(rankings, match.Red2, matchResult, true)
		}
		if !match.Red3IsSurrogate {
			addMatchResultToRankings(rankings, match.Red3, matchResult, true)
		}
		if !match.Blue1IsSurrogate {
			addMatchResultToRankings(rankings, match.Blue1, matchResult, false)
		}
		if !match.Blue2IsSurrogate {
			addMatchResultToRankings(rankings, match.Blue2, matchResult, false)
		}
		if !match.Blue3IsSurrogate {
			addMatchResultToRankings(rankings, match.Blue3, matchResult, false)
		}
	}

	// Retrieve old rankings so that we can display changes in rank as a result of this calculation.
	oldRankings, err := database.GetAllRankings()
	if err != nil {
		return nil, err
	}
	oldRankingsMap := make(map[int]game.Ranking, len(oldRankings))
	for _, ranking := range oldRankings {
		oldRankingsMap[ranking.TeamId] = ranking
	}

	sortedRankings := sortRankings(rankings)
	for rank, ranking := range sortedRankings {
		sortedRankings[rank].Rank = rank + 1
		if oldRank, ok := oldRankingsMap[ranking.TeamId]; ok {
			if preservePreviousRank {
				sortedRankings[rank].PreviousRank = oldRank.PreviousRank
			} else {
				sortedRankings[rank].PreviousRank = oldRank.Rank
			}
		}
	}
	err = database.ReplaceAllRankings(sortedRankings)
	if err != nil {
		return nil, err
	}

	return sortedRankings, nil
}

// Incrementally accounts for the given match result in the set of rankings that are being built.
func addMatchResultToRankings(
	rankings map[int]*game.Ranking, teamId int, matchResult *model.MatchResult, isRed bool,
) {
	ranking := rankings[teamId]
	if ranking == nil {
		ranking = &game.Ranking{TeamId: teamId}
		rankings[teamId] = ranking
	}

	if isRed {
		ranking.AddScoreSummary(matchResult.RedScoreSummary(), matchResult.BlueScoreSummary())
	} else {
		ranking.AddScoreSummary(matchResult.BlueScoreSummary(), matchResult.RedScoreSummary())
	}
}

func sortRankings(rankings map[int]*game.Ranking) game.Rankings {
	var sortedRankings game.Rankings
	for _, ranking := range rankings {
		sortedRankings = append(sortedRankings, *ranking)
	}
	sort.Sort(sortedRankings)
	return sortedRankings
}
