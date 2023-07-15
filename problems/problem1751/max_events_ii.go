package problem1751

import "sort"

/*
You are given an array of events where events[i] = [startDayi, endDayi, valuei].
The ith event starts at startDayi and ends at endDayi, and if you attend this event, you will receive a value of valuei.
You are also given an integer k which represents the maximum number of events you can attend.
You can only attend one event at a time. If you choose to attend an event, you must attend the entire event.
Note that the end day is inclusive: that is, you cannot attend two events where one of them starts and the other ends on the same day.
Return the maximum sum of values that you can receive by attending events.
*/

func maxValue(events [][]int, k int) int {
	var dp = make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, len(events)+1)
	}

	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] })

	for idx := len(events) - 1; idx >= 0; idx-- {
		for cnt := 1; cnt <= k; cnt++ {
			nxt := search(events, events[idx][1])
			dp[cnt][idx] = max(dp[cnt][idx+1], events[idx][2]+dp[cnt-1][nxt])
		}
	}

	return dp[k][0]
}

func search(nums [][]int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid][0] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
