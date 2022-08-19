package problem0659

/*
You are given an integer array nums that is sorted in non-decreasing order.
Determine if it is possible to split nums into one or more subsequences such that both of the following conditions are true:
Each subsequence is a consecutive increasing sequence (i.e. each integer is exactly one more than the previous integer).
All subsequences have a length of 3 or more.
Return true if you can split nums according to the above conditions, or false otherwise.
*/

func isPossible(nums []int) bool {
	var freqs = map[int]int{}
	var need = map[int]int{}
	// Caclulating frequencies
	for _, n := range nums {
		freqs[n]++
	}

	for _, n := range nums {
		if freqs[n] == 0 {
			// Already taken care of
			continue
		}
		if need[n] > 0 {
			// Needed in a sequence
			need[n]--
			freqs[n]--
			need[n+1]++
		} else if freqs[n] > 0 && freqs[n+1] > 0 && freqs[n+2] > 0 {
			// Trying to create a new sequence
			freqs[n]--
			freqs[n+1]--
			freqs[n+2]--
			need[n+3]++
		} else {
			// Doesn't fit anywhere
			return false
		}
	}
	return true
}
