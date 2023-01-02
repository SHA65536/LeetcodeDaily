package problem0520

/*
We define the usage of capitals in a word to be right when one of the following cases holds:
    All letters in this word are capitals, like "USA".
    All letters in this word are not capitals, like "leetcode".
    Only the first letter in this word is capital, like "Google".
Given a string word, return true if the usage of capitals in it is right.
*/

func detectCapitalUse(word string) bool {
	if isCap(word[0]) {
		for i := 1; i < len(word)-1; i++ {
			if isCap(word[i]) != isCap(word[i+1]) {
				return false
			}
		}
	} else {
		for i := range word {
			if isCap(word[i]) {
				return false
			}
		}
	}
	return true
}

func isCap(let byte) bool {
	return let >= 'A' && let <= 'Z'
}
