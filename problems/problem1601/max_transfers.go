package problem1601

/*
We have n buildings numbered from 0 to n - 1. Each building has a number of employees.
It's transfer season, and some employees want to change the building they reside in.
You are given an array requests where requests[i] = [fromi, toi] represents an employee's request to transfer from building fromi to building toi.
All buildings are full, so a list of requests is achievable only if for each building, the net change in employee transfers is zero.
This means the number of employees leaving is equal to the number of employees moving in.
For example if n = 3 and two employees are leaving building 0, one is leaving building 1, and one is leaving building 2,
there should be two employees moving to building 0, one employee moving to building 1, and one employee moving to building 2.
Return the maximum number of achievable requests.
*/

func maximumRequests(n int, reqs [][]int) int {
	var res int
	var diffs = make([]int, n)
	calcRequests(reqs, diffs, 0, &res)
	return res
}

func calcRequests(reqs [][]int, diffs []int, cur int, res *int) {
	if len(reqs) == 0 {
		// If we processed all the requests, check that there is no surplus
		if allZeros(diffs) {
			*res = max(*res, cur)
		}
		return
	}

	// Reject a request
	// Find the max starting from the next request
	calcRequests(reqs[1:], diffs, cur, res) // cur stays the same because we rejected

	// Accept a request
	diffs[reqs[0][0]]--
	diffs[reqs[0][1]]++
	// Find the max starting from the next request
	calcRequests(reqs[1:], diffs, cur+1, res) // cur increases because we accepted

	// Undo the request
	diffs[reqs[0][0]]++
	diffs[reqs[0][1]]--

}

func allZeros(inc []int) bool {
	for i := range inc {
		if inc[i] != 0 {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
