package problem0948

import "sort"

/*
You have an initial power of power, an initial score of 0, and a bag of tokens where tokens[i] is the value of the ith token (0-indexed).
Your goal is to maximize your total score by potentially playing each token in one of two ways:
If your current power is at least tokens[i], you may play the ith token face up, losing tokens[i] power and gaining 1 score.
If your current score is at least 1, you may play the ith token face down, gaining tokens[i] power and losing 1 score.
Each token may be played at most once and in any order. You do not have to play all the tokens.
Return the largest possible score you can achieve after playing any number of tokens.
*/

func bagOfTokensScore(tokens []int, power int) int {
	var score, maxScore int
	var start, end int = 0, len(tokens) - 1
	// Sort the tokens in increasing order
	sort.Ints(tokens)
	for start <= end {
		if power >= tokens[start] {
			// If we can 'buy' the cheapest token, do it
			power -= tokens[start]
			score++
			start++
		} else if score > 0 {
			// If we can't 'buy' the cheapest token, sell the highest one
			power += tokens[end]
			score--
			end--
		} else {
			// If we don't have enough score to sell, end
			break
		}
		// Always keep the maximum score
		maxScore = max(maxScore, score)
	}
	return maxScore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
