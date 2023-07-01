package problem2305

import "math"

/*
You are given an integer array cookies, where cookies[i] denotes the number of cookies in the ith bag.
You are also given an integer k that denotes the number of children to distribute all the bags of cookies to.
All the cookies in the same bag must go to the same child and cannot be split up.
The unfairness of a distribution is defined as the maximum total cookies obtained by a single child in the distribution.
Return the minimum unfairness of all distributions.
*/

func distributeCookies(cookies []int, k int) int {
	return calcMinUnfairness(cookies, make([]int, k))
}

// calcMinUnfairness calculates the minimum unfairness with
// cookies as cookie bags, dist as the number of cookies each child got
func calcMinUnfairness(cookies, dist []int) int {
	if len(cookies) == 0 {
		// If we ran out of cookies, find the unfairness value
		res := 0
		// Unfairness the the max cookies a child got
		for _, v := range dist {
			res = max(res, v)
		}
		return res
	}
	// If we still have cookies to give
	var res = math.MaxInt
	// Try giving the first bag to each child
	for i := range dist {
		// Give cookies
		dist[i] += cookies[0]
		// Calculate unfairness after giving the cookies
		res = min(res, calcMinUnfairness(cookies[1:], dist))
		// Take back the cookies to not interfere with the next calc
		dist[i] -= cookies[0]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
