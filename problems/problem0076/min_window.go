package problem0076

/*
Given two strings s and t of lengths m and n respectively,
return the minimum window substring of s such that every character in t (including duplicates) is included in the window.
If there is no such substring, return the empty string "".
The testcases will be generated such that the answer is unique.
A substring is a contiguous sequence of characters within the string.
*/

func minWindow(s string, t string) string {
	var need = map[byte]int{}
	var i, j, start, end, missing int
	missing = len(t)
	// Getting needed frequencies
	for i = range t {
		need[t[i]]++
	}
	// Finding window
	i = 0
	for j = 1; j <= len(s); j++ {
		char := s[j-1] // Current char we're looking at
		if need[char] > 0 {
			// If we need the current char, decrease missing amount
			missing--
		}
		need[char]-- // Update needed freq
		if missing == 0 {
			// If we have all we need, try to shrink from the left
			for i < j && need[s[i]] < 0 {
				// Shrinking as much as possible
				need[s[i]]++
				i++
			}
			need[s[i]]++
			missing++
			// Checking if this is minimal window
			if end == 0 || j-i < end-start {
				start, end = i, j
			}
			i++
		}
	}
	return s[start:end]
}
