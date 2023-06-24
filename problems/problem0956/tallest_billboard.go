package problem0956

/*
You are installing a billboard and want it to have the largest height.
The billboard will have two steel supports, one on each side. Each steel support must be an equal height.
You are given a collection of rods that can be welded together.
For example, if you have rods of lengths 1, 2, and 3, you can weld them together to make a support of length 6.
Return the largest possible height of your billboard installation. If you cannot support the billboard, return 0.
*/

func tallestBillboard(rods []int) int {
	// dp[x] represents the max height of the billboard
	// when the difference between left and right is x
	var dp = map[int]int{0: 0}
	// For each rod
	for _, rod := range rods {
		temp := map[int]int{}
		// Update the height for each difference after adding the rod
		for s := range dp {
			// If the rod is on the left
			temp[s+rod] = max(temp[s+rod], dp[s]+rod)
			// If we don't use the rod
			temp[s] = max(temp[s], dp[s])
			// If the rod is on the right
			temp[s-rod] = max(temp[s-rod], dp[s])
		}
		dp = temp
	}
	// return the height where the difference is 0
	return dp[0]
}

func tallestBillboardNaive(rods []int) int {
	return calcTallestNaive(rods, 0, 0, 0)
}

func calcTallestNaive(rods []int, left, right, idx int) int {
	var res int
	// If left and right are the same, we can make a billboard
	if left == right {
		res = max(res, left)
	}
	// If we run out of rods, return
	if idx >= len(rods) {
		return res
	}
	// Take the rod to the left
	res = max(res, calcTallestNaive(rods, left+rods[idx], right, idx+1))
	// Take the rod to the right
	res = max(res, calcTallestNaive(rods, left, right+rods[idx], idx+1))
	// Don't take the rod
	return max(res, calcTallestNaive(rods, left, right, idx+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
