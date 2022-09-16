package problem0746

/*
https://leetcode.com/problems/min-cost-climbing-stairs

You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.
*/

func minCostClimbingStairs(cost []int) int {
	var cache = map[int]int{}
	cache[len(cost)-1] = cost[len(cost)-1]
	cache[len(cost)-2] = cost[len(cost)-2]
	return min(solve(cache, cost, 0), solve(cache, cost, 1))
}

func solve(cache map[int]int, cost []int, idx int) int {
	if val, ok := cache[idx]; ok {
		return val
	}
	if cost[idx+1] > cost[idx+2] {
		cache[idx] = cost[idx] + solve(cache, cost, idx+2)
	} else {
		cache[idx] = cost[idx] + min(solve(cache, cost, idx+1), solve(cache, cost, idx+2))
	}
	return cache[idx]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
