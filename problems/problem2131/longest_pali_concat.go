package problem2131

/*
You are given an array of strings words. Each element of words consists of two lowercase English letters.
Create the longest possible palindrome by selecting some elements from words and concatenating them in any order.
Each element can be selected at most once.
Return the length of the longest palindrome that you can create. If it is impossible to create any palindrome, return 0.
A palindrome is a string that reads the same forward and backward.
*/

func longestPalindrome(words []string) int {
	var seen = map[string]int{}
	var unpaired, res int
	for i := range words {
		if words[i][0] != words[i][1] {
			rev := Reverse(words[i])
			if seen[rev] > 0 {
				res += 2
				seen[rev]--
			} else {
				seen[words[i]]++
			}
		} else {
			if seen[words[i]] > 0 {
				unpaired--
				seen[words[i]]--
				res += 2
			} else {
				seen[words[i]]++
				unpaired++
			}
		}
	}
	if unpaired > 0 {
		res += 1
	}
	return res * 2
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

const offset = 'a'

func longestPalindromeNoMap(words []string) int {
	var res int
	var seen [26][26]int
	var a, b byte
	for i := range words {
		a, b = words[i][0]-offset, words[i][1]-offset
		if seen[b][a] > 0 {
			res += 2
			seen[b][a]--
			continue
		}
		seen[a][b]++
	}
	for i := range seen {
		if seen[i][i] > 0 {
			res++
			break
		}
	}
	return res * 2
}
