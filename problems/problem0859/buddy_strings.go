package problem0859

/*
Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, otherwise, return false.
Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at s[i] and s[j].
    For example, swapping at indices 0 and 2 in "abcd" results in "cbad".
*/

func buddyStrings(s string, goal string) bool {
	// Different length - can't swap to fix
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		// If they're both the same, we can only make them equal
		// by swapping duplicate letters twice
		var set uint32 // Save existence of letters as bits
		for i := range s {
			if set&(1<<(s[i]-'a')) > 0 {
				// If we've already seen this letter, there are duplicates
				return true
			}
			// Mark letter as seen
			set |= (1 << (s[i] - 'a'))
		}
		return false
	}
	// Dif is an array of indexes where the strings differ
	var dif = make([]int, 0, len(s))
	for i := 0; i < len(s) && len(dif) <= 2; i++ {
		if s[i] != goal[i] {
			dif = append(dif, i)
		}
	}
	return len(dif) == 2 && // We can only fix two differences
		s[dif[0]] == goal[dif[1]] && // the differences should be the same letter
		s[dif[1]] == goal[dif[0]] // and should be symmetrical
}
