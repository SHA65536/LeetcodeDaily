package problem0387

/*
Given a string s, find the first non-repeating character in it and return its index.
If it does not exist, return -1.
*/

func firstUniqCharMap(s string) int {
	var freqs = map[byte]int{}
	for i := range s {
		freqs[s[i]]++
	}
	for i := range s {
		if freqs[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar(s string) int {
	var freqs [26]int
	for i := range s {
		freqs[s[i]-'a']++
	}
	for i := range s {
		if freqs[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
