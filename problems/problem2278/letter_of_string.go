package problem2278

/*
Given a string s and a character letter, return the percentage of characters in s that equal letter rounded down to the nearest whole percent.
*/

func percentageLetter(s string, letter byte) int {
	var cnt int
	for i := range s {
		if s[i] == letter {
			cnt++
		}
	}
	return (cnt * 100) / len(s)
}
