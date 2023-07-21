package problem0673

/*
Given an integer array nums, return the number of longest increasing subsequences.
Notice that the sequence has to be strictly increasing.
*/

func findNumberOfLIS(nums []int) int {
	var res, maxlen int
	// length[i] is the length of the longest subseq that ends at position i
	var length = make([]int, len(nums))
	// count[i] is the number of longest subseqs that end at position i
	var count = make([]int, len(nums))

	// For each ending position
	for end := range nums {
		length[end], count[end] = 1, 1 // the smallest subseq length is 1
		// Try find a previous ending position, that is smaller
		for prev := 0; prev < end; prev++ {
			// If the previous subseq can be continued,
			// Update the length and count of the subseqs
			if nums[end] > nums[prev] {
				if length[end] == length[prev]+1 {
					count[end] += count[prev]
				} else if length[end] < length[prev]+1 {
					length[end] = length[prev] + 1
					count[end] = count[prev]
				}
			}
		}
	}

	// Sum up the counts of the subsequences with the max length
	for i := range length {
		if length[i] > maxlen {
			maxlen = length[i]
			res = count[i]
		} else if length[i] == maxlen {
			res += count[i]
		}
	}

	return res
}
