package problem1569

/*
Given an array nums that represents a permutation of integers from 1 to n.
We are going to construct a binary search tree (BST) by inserting the elements of nums in order into an initially empty BST.
Find the number of different ways to reorder nums so that the constructed BST is identical to that formed from the original array nums.
    For example, given nums = [2,1,3], we will have 2 as the root, 1 as a left child, and 3 as a right child.
	The array [2,3,1] also yields the same BST but [3,2,1] yields a different BST.
Return the number of ways to reorder nums such that the BST formed is identical to the original BST formed from nums.
Since the answer may be very large, return it modulo 109 + 7.
*/

const mod int64 = 1e9 + 7

func numOfWays(nums []int) int {
	var cache = make([][]int64, len(nums)+1)

	for i := 0; i <= len(nums); i++ {
		cache[i] = fill(len(nums)+1, 1)
		for j := 1; j < i; j++ {
			cache[i][j] = (cache[i-1][j-1] + cache[i-1][j]) % mod
		}
	}

	return int(dfs(nums, cache)%mod - 1)
}

func dfs(nums []int, cache [][]int64) int64 {
	if len(nums) <= 2 {
		return 1
	}

	var left = make([]int, 0, len(nums))
	var right = make([]int, 0, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	var leftRes = dfs(left, cache) % mod
	var rightRes = dfs(right, cache) % mod

	return (((cache[len(nums)-1][len(left)] * leftRes) % mod) * rightRes) % mod
}

func fill(size int, val int64) []int64 {
	var res = make([]int64, size)
	for i := range res {
		res[i] = val
	}
	return res
}
