package problem1335

/*
You want to schedule a list of jobs in d days.
Jobs are dependent (i.e To work on the ith job, you have to finish all the jobs j where 0 <= j < i).
You have to finish at least one task every day. The difficulty of a job schedule is the sum of difficulties of each day of the d days.
The difficulty of a day is the maximum difficulty of a job done on that day.
You are given an integer array jobDifficulty and an integer d. The difficulty of the ith job is jobDifficulty[i].
Return the minimum difficulty of a job schedule. If you cannot find a schedule for the jobs return -1.
*/

const MAX_INT = 1e9

func minDifficulty(jobDifficulty []int, d int) int {
	var cache [][]int
	if len(jobDifficulty) < d {
		return -1
	}
	cache = make([][]int, len(jobDifficulty))
	for i := range cache {
		cache[i] = make([]int, d+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return solveMinDiff(d, 0, jobDifficulty, cache)
}

func solveMinDiff(d, l int, jobs []int, cache [][]int) int {
	var cmax, cmin, temp int
	if d == 0 && l == len(jobs) {
		return 0
	}
	if d == 0 || l == len(jobs) {
		return MAX_INT
	}
	if cache[l][d] != -1 {
		return cache[l][d]
	}

	cmax = jobs[l]
	cmin = MAX_INT
	for i := l; i < len(jobs); i++ {
		if cmax < jobs[i] {
			cmax = jobs[i]
		}
		temp = solveMinDiff(d-1, i+1, jobs, cache)
		if temp != MAX_INT && cmin > temp+cmax {
			cmin = temp + cmax
		}
	}
	cache[l][d] = cmin
	return cmin
}

func minDifficulty1D(jobDifficulty []int, d int) int {
	var max_dif int
	var dp []int
	if len(jobDifficulty) < d {
		return -1
	}

	dp = make([]int, len(jobDifficulty)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = MAX_INT
	}
	dp[len(jobDifficulty)] = 0

	for day := 1; day <= d; day++ {
		for i := 0; i <= len(jobDifficulty)-day; i++ {
			max_dif, dp[i] = 0, MAX_INT
			for j := i; j <= len(jobDifficulty)-day; j++ {
				if max_dif < jobDifficulty[j] {
					max_dif = jobDifficulty[j]
				}
				if dp[i] > max_dif+dp[j+1] {
					dp[i] = max_dif + dp[j+1]
				}
			}
		}
	}
	return dp[0]
}
