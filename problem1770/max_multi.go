package problem1770

/*
You are given two integer arrays nums and multipliers of size n and m respectively, where n >= m. The arrays are 1-indexed.
You begin with a score of 0. You want to perform exactly m operations. On the ith operation (1-indexed), you will:
Choose one integer x from either the start or the end of the array nums.
Add multipliers[i] * x to your score.
Remove x from the array nums.
Return the maximum score after performing m operations.
*/

func maximumScore(nums []int, multipliers []int) int {
	// Cache to look up previously calculated states
	var cache = map[[2]int]int{}
	return maxScore(cache, nums, multipliers, 0, len(nums)-1, 0)
}

func maxScore(cache map[[2]int]int, nums, multi []int, left, right, turn int) int {
	// If were at the end, return 0
	if turn == len(multi) {
		return 0
	}
	// If we calculated this state, return the cache
	if val, ok := cache[[2]int{left, right}]; ok {
		return val
	}
	// Taking a number from the left
	sumLeft := nums[left]*multi[turn] + maxScore(cache, nums, multi, left+1, right, turn+1)
	// Taking a number from the right
	sumRight := nums[right]*multi[turn] + maxScore(cache, nums, multi, left, right-1, turn+1)
	// Picking the higher one and adding to the cache
	if sumLeft > sumRight {
		cache[[2]int{left, right}] = sumLeft
	} else {
		cache[[2]int{left, right}] = sumRight
	}
	return cache[[2]int{left, right}]
}
