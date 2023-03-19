package problem0912

/*
Given an array of integers nums, sort the array in ascending order and return it.
You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.
*/

// This is a merge sort
func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	// Split the array in 2
	a, b := nums[:len(nums)/2], nums[len(nums)/2:]
	// Sort each array recursively
	a = sortArray(a)
	b = sortArray(b)
	// Merge both arrays together
	return mergeArray(a, b)
}

// Merge two sorted arrays together
func mergeArray(a, b []int) []int {
	var res = make([]int, 0, len(a)+len(b))
	// Loop through both arrays
	var aidx, bidx int
	for aidx < len(a) && bidx < len(b) {
		// Pick the lower array
		if a[aidx] < b[bidx] {
			// Add to res
			res = append(res, a[aidx])
			aidx++
		} else {
			res = append(res, b[bidx])
			bidx++
		}
	}
	// Take the slack from the array if there is any
	for aidx < len(a) {
		res = append(res, a[aidx])
		aidx++
	}
	for bidx < len(b) {
		res = append(res, b[bidx])
		bidx++
	}
	return res
}

func sortArrayQuick(nums []int) []int {
	return quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, low, high int) []int {
	if low < high {
		var p int
		nums, p = partition(nums, low, high)
		nums = quickSort(nums, low, p-1)
		nums = quickSort(nums, p+1, high)
	}
	return nums
}

func partition(nums []int, low, high int) ([]int, int) {
	pIdx := findPivot(nums, low, high)
	nums[high], nums[pIdx] = nums[pIdx], nums[high]
	pivot := nums[high]
	i := low
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[high] = nums[high], nums[i]
	return nums, i
}

// findPivot finds the median of the start, middle, and end of the array
func findPivot(nums []int, low, high int) int {
	a, b, c := nums[low], nums[(low+high)/2], nums[high]
	if (b >= a) == (a >= c) {
		return low
	} else if (a >= b) == (b >= c) {
		return (low + high) / 2
	} else if (a >= c) == (c >= b) {
		return high
	}
	return low
}
