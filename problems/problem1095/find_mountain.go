package problem1095

/*
You may recall that an array arr is a mountain array if and only if:
arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
Given a mountain array mountainArr, return the minimum index such that mountainArr.get(index) == target.
If such an index does not exist, return -1.
You cannot access the mountain array directly. You may only access the array using a MountainArray interface:
MountainArray.get(k) returns the element of the array at index k (0-indexed).
MountainArray.length() returns the length of the array.
Submissions making more than 100 calls to MountainArray.get will be judged Wrong Answer.
Also, any solutions that attempt to circumvent the judge will result in disqualification.
*/

func findInMountainArray(target int, m *MountainArray) int {
	// Find the peak of the mountain
	var pIdx = findPeak(m)

	// Search the left side first
	var leftIdx = binarySearch(0, pIdx, target, m, small)
	if m.get(leftIdx) == target {
		return leftIdx
	}

	// If not in the left, search the right
	var rightIdx = binarySearch(pIdx+1, m.length()-1, target, m, big)
	if m.get(rightIdx) == target {
		return rightIdx
	}

	return -1
}

// We can use binary search to find the peak
func findPeak(m *MountainArray) int {
	var low, high = 1, m.length() - 2
	for low < high {
		var mid = low + (high-low)/2
		if m.get(mid) < m.get(mid+1) { // If increasing, look at right side
			low = mid + 1
		} else { // If decreasing, look at left side
			high = mid
		}
	}
	return low
}

type comp func(int, int) bool

func big(a, b int) bool   { return a > b }
func small(a, b int) bool { return a < b }

// Generalized binary search with dynamic compare func
func binarySearch(low, high, target int, m *MountainArray, c comp) int {
	for low < high {
		var mid = low + (high-low)/2
		if c(m.get(mid), target) {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
