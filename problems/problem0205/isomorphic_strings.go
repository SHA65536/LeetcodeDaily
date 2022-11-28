package problem0205

/*
Given two strings s and t, determine if they are isomorphic.
Two strings s and t are isomorphic if the characters in s can be replaced to get t.
All occurrences of a character must be replaced with another character while preserving the order of characters.
No two characters may map to the same character, but a character may map to itself.
*/

func isIsomorphic(s string, t string) bool {
	var sm, tm [256]int //sm and tm represent the last time we've seen i in s and in t
	for i := range s {
		// If the last time we've seen them is not the same
		// it means they don't map to the same char
		if sm[s[i]] != tm[t[i]] {
			return false
		}
		// Updating last time we've seen them
		sm[s[i]] = i + 1 // (+ 1 is just to make it work with 0 indexed arrays)
		tm[t[i]] = i + 1
	}
	return true
}
