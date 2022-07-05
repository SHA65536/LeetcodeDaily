package problem0128

/*
Given an unsorted array of integers nums,
return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.
*/

func longestConsecutive(nums []int) int {
	var max int
	var streaks = map[int]bool{}
	// Indexing the numbers and removing duplicates
	for _, num := range nums {
		streaks[num] = true
	}
	// Calculating highest streak
	for key := range streaks {
		// Ignoring keys in the middle of a streak
		if _, ok := streaks[key-1]; !ok {
			var cur int
			// Checking continuation
			for cur = key + 1; streaks[cur]; cur++ {
			}
			if cur-key > max {
				max = cur - key
			}
		}
	}
	return max
}
