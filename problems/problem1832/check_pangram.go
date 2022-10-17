package problem1832

/*
A pangram is a sentence where every letter of the English alphabet appears at least once.
Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.
*/

const OFFSET = 'a'

func checkIfPangram(sentence string) bool {
	var u_freq int
	var letters [26]bool
	if len(sentence) < 26 {
		return false
	}
	for i := 0; i < len(sentence) && u_freq < 26; i++ {
		if !letters[sentence[i]-OFFSET] {
			letters[sentence[i]-OFFSET] = true
			u_freq++
		}
	}
	return u_freq == 26
}
