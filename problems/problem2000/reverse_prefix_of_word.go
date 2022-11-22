package problem2000

import "strings"

/*
Given a 0-indexed string word and a character ch,
reverse the segment of word that starts at index 0 and ends at the index of the first occurrence of ch (inclusive).
If the character ch does not exist in word, do nothing.
For example, if word = "abcdefd" and ch = "d", then you should reverse the segment that starts at 0 and ends at 3 (inclusive).
The resulting string will be "dcbaefd".
Return the resulting string.
*/

func reversePrefix(word string, ch byte) string {
	var end int = strings.IndexByte(word, ch)
	var data []byte
	if end == -1 {
		return word
	}
	data = []byte(word)
	for i := 0; end > i; i++ {
		data[i], data[end] = data[end], data[i]
		end--
	}
	return string(data)
}
