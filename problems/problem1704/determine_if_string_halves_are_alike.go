package problem1704

/*
You are given a string s of even length. Split this string into two halves of equal lengths, and let a be the first half and b be the second half.
Two strings are alike if they have the same number of vowels ('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U').
Notice that s contains uppercase and lowercase letters.
Return true if a and b are alike. Otherwise, return false.
*/

var VowelMap = map[byte]bool{
	'A': true, 'E': true, 'I': true,
	'O': true, 'U': true,
}

func halvesAreAlike(s string) bool {
	var vowels int
	for i := 0; i < len(s)/2; i++ {
		if VowelMap[s[i]] || VowelMap[s[i]-32] {
			vowels++
		}
	}
	for i := len(s) / 2; i < len(s) && vowels >= 0; i++ {
		if VowelMap[s[i]] || VowelMap[s[i]-32] {
			vowels--
		}
	}
	return vowels == 0
}
