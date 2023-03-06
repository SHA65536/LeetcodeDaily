package problem1539

/*
Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.
Return the kth positive integer that is missing from this array.
*/

func findKthPositive(arr []int, k int) int {
	// Finding the missing number using binary search
	var left, right = 0, len(arr)
	for right > left {
		mid := (left + right) / 2
		// arr[mid]-mid-1 is the number of missing numbers left of mid
		if arr[mid]-mid-1 < k {
			// If there are less than k numbers missing, look at the right half
			left = mid + 1
		} else {
			// If ther are more than knumbers missing, look at the left half
			right = mid
		}
	}
	return right + k
}
