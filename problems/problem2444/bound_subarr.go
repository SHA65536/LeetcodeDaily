package problem2444

/*
You are given an integer array nums and two integers minK and maxK.
A fixed-bound subarray of nums is a subarray that satisfies the following conditions:
    The minimum value in the subarray is equal to minK.
    The maximum value in the subarray is equal to maxK.
Return the number of fixed-bound subarrays.
A subarray is a contiguous part of an array.
*/

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64
	var start, lastMin, lastMax int = 0, -1, -1
	for i := range nums {
		if nums[i] < minK || nums[i] > maxK {
			lastMin, lastMax = -1, -1
			start = i + 1
		}
		if nums[i] == minK {
			lastMin = i
		}
		if nums[i] == maxK {
			lastMax = i
		}
		res += int64(max(0, min(lastMin, lastMax)-start+1))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
