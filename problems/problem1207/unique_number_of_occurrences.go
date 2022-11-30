package problem1207

/*
Given an array of integers arr, return true if the number of occurrences of each value in the array is unique, or false otherwise.
*/

func uniqueOccurrences(arr []int) bool {
	var freq = map[int]int{}
	var dfreq = map[int]bool{}
	for i := range arr {
		freq[arr[i]]++
	}
	for i := range freq {
		if dfreq[freq[i]] {
			return false
		}
		dfreq[freq[i]] = true
	}
	return true
}
