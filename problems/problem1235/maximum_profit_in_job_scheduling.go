package problem1235

import "sort"

/*
We have n jobs, where every job is scheduled to be done from startTime[i] to endTime[i], obtaining a profit of profit[i].
You're given the startTime, endTime and profit arrays, return the maximum profit you can take such that there are no two jobs
in the subset with overlapping time range.
If you choose a job that ends at time X you will be able to start another job that starts at time X.
*/

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var jobs [][3]int
	var memo = map[int]int{}
	var dp func(int) int

	// Building one array from the three inputs
	jobs = make([][3]int, len(profit)) // [start, end, profit]
	for i := range startTime {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	// Sorting by start time
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	// dp[i] is max profit we can make starting at time i
	dp = func(i int) int {
		var ifSkip, ifTake int
		var left, right, mid int
		if i == len(startTime) {
			// If we reached the end, no profit
			return 0
		}
		if v, ok := memo[i]; ok {
			// If we already calculated this, return cached
			return v
		}

		// Profit if we skip this job
		ifSkip = dp(i + 1)

		// Finding next job with binary search
		left, right = i+1, len(profit)
		for left < right {
			mid = left + (right-left)/2
			if jobs[mid][0] < jobs[i][1] {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// Profit if we take this job
		ifTake = jobs[i][2] + dp(left)

		// Cache to avoid recalculations
		memo[i] = max(ifTake, ifSkip)
		return memo[i]
	}
	return dp(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
