package problem0091

/*
A message containing letters from A-Z can be encoded into numbers using the following mapping:
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above
(there may be multiple ways). For example, "11106" can be mapped into:
"AAJF" with the grouping (1 1 10 6)
"KJF" with the grouping (11 10 6
Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".
Given a string s containing only digits, return the number of ways to decode it.
The test cases are generated so that the answer fits in a 32-bit integer.
*/

func numDecodings(s string) int {
	var one, two byte
	// cache[i] will represent the number of decodings up to character i
	var cache = make([]int, len(s))
	// starting with 0 is invalid
	if s[0] == '0' {
		return 0
	}
	cache[0] = 1
	for i := 1; i < len(s); i++ {
		one = s[i] - '0'            // Current letter
		two = (s[i-1]-'0')*10 + one // Current 2 letters
		if one > 0 {
			// If current letter is valid (not 0)
			// then path is valid
			cache[i] += cache[i-1]
		}
		if two >= 10 && two <= 26 {
			// If current 2 letters are valid
			// then path is valid
			if i > 1 {
				cache[i] += cache[i-2]
			} else {
				cache[i] += 1
			}
		}
	}
	return cache[len(s)-1]
}
