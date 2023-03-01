package problem0044

/*
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:
    '?' Matches any single character.
    '*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).
*/

func isMatch(str, pat string) bool {
	var s, p, match, star int
	star = -1

	for s < len(str) {
		if p < len(pat) && (pat[p] == '?' || str[s] == pat[p]) {
			s++
			p++
		} else if p < len(pat) && pat[p] == '*' {
			star = p
			match = s
			p++
		} else if star != -1 {
			p = star + 1
			match++
			s = match
		} else {
			return false
		}
	}
	for p < len(pat) && pat[p] == '*' {
		p++
	}

	return p == len(pat)
}
