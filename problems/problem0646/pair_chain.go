package problem0646

import (
	"sort"
)

/*
You are given an array of n pairs pairs where pairs[i] = [lefti, righti] and lefti < righti.
A pair p2 = [c, d] follows a pair p1 = [a, b] if b < c. A chain of pairs can be formed in this fashion.
Return the length longest chain which can be formed.
You do not need to use up all the given intervals. You can select pairs in any order.
*/

func findLongestChain(pairs [][]int) int {
	var res int
	var cur = -1001 // problem constraint
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	for _, p := range pairs {
		if p[0] > cur {
			cur = p[1]
			res++
		}
	}
	return res
}
