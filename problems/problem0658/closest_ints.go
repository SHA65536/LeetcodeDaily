package problem0658

/*
Given a sorted integer array arr, two integers k and x, return the k closest integers to x in the array.
The result should also be sorted in ascending order.
An integer a is closer to x than an integer b if:
|a - x| < |b - x|, or
|a - x| == |b - x| and a < b
*/

func findClosestElements(arr []int, k int, x int) []int {
	var left, right int
	right = searchInsert(arr, x)
	left = right - 1
	for i := 0; i < k; i++ {
		if left < 0 {
			right++
		} else if right >= len(arr) {
			left--
		} else if diff(x, arr[left]) > diff(x, arr[right]) {
			right++
		} else {
			left--
		}
	}
	return arr[left+1 : right]
}

// Binary search
func searchInsert(nums []int, target int) int {
	var res, start, end = -1, 0, len(nums) - 1
	for end >= start {
		mid := ((end - start) / 2) + start
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			end = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid + 1
			}
			start = mid + 1
		}
	}
	return res
}

func diff(a, b int) int {
	a = a - b
	if a < 0 {
		return -a
	}
	return a
}
