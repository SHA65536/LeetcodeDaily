package problem2448

/*
You are given two 0-indexed arrays nums and cost consisting each of n positive integers.
You can do the following operation any number of times:
    Increase or decrease any element of the array nums by 1.
The cost of doing one operation on the ith element is cost[i].
Return the minimum total cost such that all the elements of the array nums become equal.
*/

func minCost(nums []int, cost []int) int64 {
	var low, high = nums[0], nums[0]
	for i := range nums {
		low = min(low, nums[i])
		high = max(high, nums[i])
	}
	for low < high {
		mid := (low + high) / 2
		l, r := CalcCost(nums, cost, mid), CalcCost(nums, cost, mid+1)
		if l > r {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return CalcCost(nums, cost, low)
}

func CalcCost(nums, cost []int, val int) int64 {
	var res int64

	for i := range nums {
		res += int64(abs(nums[i]-val) * cost[i])
	}

	return res
}

func minCostMap(nums []int, cost []int) int64 {
	var cache = make(map[int]int64)
	var low, high = nums[0], nums[0]
	for i := range nums {
		low = min(low, nums[i])
		high = max(high, nums[i])
	}
	for low < high {
		mid := (low + high) / 2
		l, r := CalcCostMap(cache, nums, cost, mid), CalcCostMap(cache, nums, cost, mid+1)
		if l > r {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return CalcCost(nums, cost, low)
}

func CalcCostMap(cache map[int]int64, nums, cost []int, val int) int64 {
	var res int64

	if v, ok := cache[val]; ok {
		return v
	}

	for i := range nums {
		res += int64(abs(nums[i]-val) * cost[i])
	}

	cache[val] = res

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
