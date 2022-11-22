package problem1417

/*
You are given an alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).
You have to find a permutation of the string where no letter is followed by another letter and no digit is followed by another digit.
That is, no two adjacent characters have the same type.
Return the reformatted string or return an empty string if it is impossible to reformat the string.
*/

func reformat(s string) string {
	var result []byte
	var nums, chars int
	var total int = len(s)
	// Counting total
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			nums++
		} else {
			chars++
		}
		// If we can't make a good string, exit early
		if chars > (total/2)+1 || nums > (total/2)+1 {
			return ""
		}
	}
	if chars-nums >= 2 || nums-chars >= 2 {
		return ""
	}
	result = make([]byte, total)
	// The bigger ones should start
	if chars > nums {
		chars, nums = 0, 1 // Chars starts at 0, numbers at 1
	} else {
		chars, nums = 1, 0 // Chars starts at 1, numbers at 0
	}
	// Writing the new string
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			result[nums] = s[i]
			nums += 2 // A number every 2nd position
		} else {
			result[chars] = s[i]
			chars += 2 // A char every 2nd position
		}
	}
	return string(result)
}
