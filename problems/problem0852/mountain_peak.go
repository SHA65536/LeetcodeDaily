package problem0852

/*
An array arr a mountain if the following properties hold:
    - arr.length >= 3
    - There exists some i with 0 < i < arr.length - 1 such that:
        arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
        arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
Given a mountain array arr,
return the index i such that arr[0] < arr[1] < ... < arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].
You must solve it in O(log(arr.length)) time complexity.
*/

func peakIndexInMountainArray(arr []int) int {
	var low, high = 1, len(arr) - 1
	// Simple binary search
	for low < high {
		mid := low + (high-low)/2
		// Instead of looking at the value, we look at the difference
		if arr[mid]-arr[mid-1] > 0 {
			// Gradient to the right
			low = mid + 1
		} else {
			// Gradient to the left
			high = mid
		}
	}
	return low - 1
}
