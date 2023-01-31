package problem1626

import "sort"

/*
You are the manager of a basketball team. For the upcoming tournament, you want to choose the team with the highest overall score.
The score of the team is the sum of scores of all the players in the team.
However, the basketball team is not allowed to have conflicts.
A conflict exists if a younger player has a strictly higher score than an older player.
A conflict does not occur between players of the same age.
Given two lists, scores and ages, where each scores[i] and ages[i] represents the score and age of the ith player,
respectively, return the highest overall score of all possible basketball teams.
*/

func bestTeamScore(scores []int, ages []int) int {
	var res int
	// players is the order of the players for sorting by age
	var players []int = make([]int, len(scores))
	// dp[i] represents the score if the i player is included
	var dp []int = make([]int, len(scores))
	for i := range players {
		players[i] = i
	}
	// Sorting players by age
	sort.Slice(players, func(i, j int) bool {
		if ages[players[i]] == ages[players[j]] {
			return scores[players[i]] < scores[players[j]]
		}
		return ages[players[i]] < ages[players[j]]
	})
	for i := range players {
		// Adding the first player to the team
		dp[i] = scores[players[i]]
		// Checking all younger players
		for j := 0; j < i; j++ {
			// If the younger player does not have a conflict
			if scores[players[j]] <= scores[players[i]] {
				// Add him to the team if it's better than before
				dp[i] = max(dp[i], dp[j]+scores[players[i]])
			}
		}
		// Keep track of the best score
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
