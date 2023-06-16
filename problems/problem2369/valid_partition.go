package problem2369

/*
You are given a 0-indexed integer array nums. You have to partition the array into one or more contiguous subarrays.

We call a partition of the array valid if each of the obtained subarrays satisfies one of the following conditions:
    The subarray consists of exactly 2 equal elements. For example, the subarray [2,2] is good.
    The subarray consists of exactly 3 equal elements. For example, the subarray [4,4,4] is good.
    The subarray consists of exactly 3 consecutive increasing elements, that is, the difference between adjacent elements is 1.
	For example, the subarray [3,4,5] is good, but the subarray [1,3,5] is not.

Return true if the array has at least one valid partition. Otherwise, return false.
*/

func validPartition(nums []int) bool {
	var cache = make([]*bool, len(nums))
	return checkPartition(nums, cache, 0)
}

func checkPartition(nums []int, cache []*bool, idx int) bool {
	// If we reached the end, we are done
	if idx == len(nums) {
		return true
	}
	// If we already calculated this
	if cache[idx] != nil {
		return *cache[idx]
	}

	// Exactly two equal
	if idx < len(nums)-1 && nums[idx] == nums[idx+1] {
		cache[idx] = newBool(checkPartition(nums, cache, idx+2))
		if *cache[idx] {
			return true
		}
	}

	// Exactly three equal
	if idx < len(nums)-2 && nums[idx] == nums[idx+1] && nums[idx] == nums[idx+2] {
		cache[idx] = newBool(checkPartition(nums, cache, idx+3))
		if *cache[idx] {
			return true
		}
	}

	// Exactly three consecutive
	if idx < len(nums)-2 && nums[idx] == nums[idx+1]-1 && nums[idx] == nums[idx+2]-2 {
		cache[idx] = newBool(checkPartition(nums, cache, idx+3))
		if *cache[idx] {
			return true
		}
	}

	cache[idx] = newBool(false)
	return false
}

func newBool(val bool) *bool {
	var res = new(bool)
	*res = val
	return res
}
