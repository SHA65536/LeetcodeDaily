package problem2129

import "strings"

/*
You are given a string title consisting of one or more words separated by a single space, where each word consists of English letters.
Capitalize the string by changing the capitalization of each word such that:
    If the length of the word is 1 or 2 letters, change all letters to lowercase.
    Otherwise, change the first letter to uppercase and the remaining letters to lowercase.
Return the capitalized title.
*/

func capitalizeTitle(title string) string {
	var res strings.Builder
	var words = strings.Split(title, " ")
	for i := range words {
		res.WriteByte(' ')
		if len(words[i]) < 3 {
			res.WriteString(strings.ToLower(words[i]))
		} else {
			if isCap(words[i][0]) {
				res.WriteByte(words[i][0])
			} else {
				res.WriteByte(words[i][0] - 32)
			}
			res.WriteString(strings.ToLower(words[i][1:]))
		}
	}
	return res.String()[1:]
}

func isCap(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func capitalizeTitleBit(title string) string {
	var result = make([]byte, len(title))
	index := 0
	for i := 0; i < len(title); i++ {
		if title[i] == ' ' || i == (len(title)-1) {
			if i-index > 2 || (i == (len(title)-1) && i-index >= 2) {
				result[index] &= '_'
			}
			index = i + 1
		}
		result[i] = title[i] | ' '
	}
	return string(result)
}
