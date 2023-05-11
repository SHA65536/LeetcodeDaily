package problem1035

/*
You are given two integer arrays nums1 and nums2.
We write the integers of nums1 and nums2 (in the order they are given) on two separate horizontal lines.
We may draw connecting lines: a straight line connecting two numbers nums1[i] and nums2[j] such that:
    nums1[i] == nums2[j], and
    the line we draw does not intersect any other connecting (non-horizontal) line.
Note that a connecting line cannot intersect even at the endpoints (i.e., each number can only belong to one connecting line).
Return the maximum number of connecting lines we can draw in this way.
*/

func maxUncrossedLines(nums1, nums2 []int) int {
	// dp[i][j] is the longest common subsequnce up to
	// nums1[i] and nums2[j]
	var dp = make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				// If we found a match, increase the next cell
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// If no match, take the max of the previous 2 cells
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	// dp[len(nums1)][len(nums2)] is the longest common subsequence
	// up to the end of the inputs
	return dp[len(nums1)][len(nums2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
