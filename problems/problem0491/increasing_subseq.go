package problem0491

/*
Given an integer array nums,
return all the different possible non-decreasing subsequences of the given array with at least two elements.
You may return the answer in any order.
*/

func findSubsequences(nums []int) [][]int {
	var res [][]int
	// cur contains the sub sequence we're currently looking at
	var cur []int
	var dfs func(pos int)
	dfs = func(pos int) {
		if len(cur) > 1 {
			// Copying the current subsequence to res
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
		}
		var set = map[int]bool{}
		// Trying to add each number to the current sequence
		for i := pos; i < len(nums); i++ {
			// If it's non-decreasing and we haven't tried to add the same number already
			if (len(cur) == 0 || nums[i] >= cur[len(cur)-1]) && !set[nums[i]] {
				// Add to current subsequence
				cur = append(cur, nums[i])
				// Explore all subsequences from there
				dfs(i + 1)
				// Remove from current subsequence
				cur = cur[:len(cur)-1]
				set[nums[i]] = true
			}
		}
	}
	dfs(0)
	return res
}
