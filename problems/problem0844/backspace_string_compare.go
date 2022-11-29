package problem0844

/*
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.
Note that after backspacing an empty text, the text will continue empty.
*/

func backspaceCompare(s string, t string) bool {
	var sIdx, tIdx int = len(s) - 1, len(t) - 1
	var toDelete int
	for sIdx >= 0 || tIdx >= 0 {
		// If we need to delete from s
		for sIdx >= 0 && s[sIdx] == '#' {
			toDelete = 1
			sIdx--
			// Count the number of chars to delete
			for sIdx >= 0 && toDelete > 0 {
				if s[sIdx] == '#' {
					toDelete++
				} else {
					toDelete--
				}
				sIdx--
			}
		}
		// If we need to delete from t
		for tIdx >= 0 && t[tIdx] == '#' {
			toDelete = 1
			tIdx--
			// Count the number of chars to delete
			for tIdx >= 0 && toDelete > 0 {
				if t[tIdx] == '#' {
					toDelete++
				} else {
					toDelete--
				}
				tIdx--
			}
		}
		// If one is fully deleted and the other is not
		if (sIdx >= 0 && tIdx < 0) || (tIdx >= 0 && sIdx < 0) {
			return false
		}
		// If the characters don't match
		if sIdx >= 0 && tIdx >= 0 && s[sIdx] != t[tIdx] {
			return false
		}
		sIdx--
		tIdx--
	}
	// If both strings finished return true
	return sIdx <= 0 && tIdx <= 0
}
