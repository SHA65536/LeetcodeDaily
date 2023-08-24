package problem0767

/*
Given a string s, rearrange the characters of s so that any two adjacent characters are not the same.
Return any possible rearrangement of s or return "" if not possible.
*/

func reorganizeString(s string) string {
	var res = make([]byte, len(s))

	// Calculate frequencies
	var freq [26]int
	for i := range s {
		freq[s[i]-'a']++
	}

	// Find the max
	var max, mIdx int
	for i, v := range freq {
		if v > max {
			max = v
			mIdx = i
		}
	}

	// Can't reorganize string if max is more than half
	// (no space for the most common letter)
	if max > (len(s)+1)/2 {
		return ""
	}

	// Distribute the max letter in even positions
	var idx int
	for freq[mIdx] > 0 {
		res[idx] = byte(mIdx) + 'a'
		idx += 2
		freq[mIdx]--
	}

	// Distribute remaining letters
	for i := range freq {
		for freq[i] > 0 {
			// If we reaced the end
			if idx >= len(s) {
				// Start again in odd positions
				idx = 1
			}
			res[idx] = byte(i) + 'a'
			idx += 2
			freq[i]--
		}
	}

	return string(res)
}
