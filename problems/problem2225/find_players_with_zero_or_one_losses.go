package problem2225

import "sort"

/*
You are given an integer array matches where matches[i] = [winneri, loseri] indicates that the player winneri defeated player loseri in a match.
Return a list answer of size 2 where:
answer[0] is a list of all players that have not lost any matches.
answer[1] is a list of all players that have lost exactly one match.
The values in the two lists should be returned in increasing order.

Note:
You should only consider the players that have played at least one match.
The testcases will be generated such that no two matches will have the same outcome.
*/

func findWinners(matches [][]int) [][]int {
	var results = [][]int{{}, {}}
	var losses = make(map[int]int)
	for i := range matches {
		if _, ok := losses[matches[i][0]]; !ok {
			losses[matches[i][0]] = 0
		}
		losses[matches[i][1]]++
	}
	for k, v := range losses {
		if v == 0 {
			results[0] = append(results[0], k)
		} else if v == 1 {
			results[1] = append(results[1], k)
		}
	}
	sort.Ints(results[0])
	sort.Ints(results[1])
	return results
}
