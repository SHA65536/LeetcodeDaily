package problem0880

/*
You are given an encoded string s.
To decode the string to a tape, the encoded string is read one character at a time and the following steps are taken:
- If the character read is a letter, that letter is written onto the tape.
- If the character read is a digit d, the entire current tape is repeatedly written d - 1 more times in total.
Given an integer k, return the kth letter (1-indexed) in the decoded string.
*/

func decodeAtIndex(s string, k int) string {
	var idx, ln, rk int64
	rk = int64(k)

	// Find total length
	for ln < rk {
		if isDigit(s[idx]) {
			ln *= int64(s[idx] - '0')
		} else {
			ln++
		}
		idx++
	}

	// Iterating in reverse
	for i := idx - 1; i >= 0; i-- {
		if isDigit(s[i]) {
			ln /= int64(s[i] - '0')
			rk %= ln
		} else {
			if rk == 0 || rk == ln {
				return s[i : i+1]
			}
			ln--
		}
	}

	return ""
}

func isDigit(a byte) bool {
	return a >= '0' && a <= '9'
}
