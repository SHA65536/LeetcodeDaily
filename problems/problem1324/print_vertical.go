package problem1324

import "strings"

/*
Given a string s. Return all the words vertically in the same order in which they appear in s.
Words are returned as a list of strings, complete with spaces when is necessary. (Trailing spaces are not allowed).
Each word would be put on only one column and that in one column there will be only one word.
*/

func printVertically(s string) []string {
	// res is words rotated
	var res [][]byte
	// done is res converted to strings
	var done []string
	// last is the position of the last character of each row
	var last []int
	var longest int
	words := strings.Split(s, " ")
	// Finding longest word
	for i := range words {
		if len(words[i]) > longest {
			longest = len(words[i])
		}
	}
	// There should be as many rows as the number of chars in the longest word
	res = make([][]byte, longest)
	done = make([]string, longest)
	last = make([]int, longest)
	// For each row
	for i := 0; i < longest; i++ {
		var max int
		res[i] = make([]byte, len(words))
		// Get the letter from each word
		for j := range words {
			if i < len(words[j]) {
				// If the word is long enough, put a letter
				res[i][j] = words[j][i]
				max = j
			} else {
				// If it's not long enough, put a space
				res[i][j] = ' '
			}
		}
		last[i] = max
	}
	// Convert to strings and remove trailing spaces
	for i := range res {
		done[i] = string(res[i][:last[i]+1])
	}
	return done
}
