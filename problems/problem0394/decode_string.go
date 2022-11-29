package problem0394

import "strings"

/*
Given an encoded string, return its decoded string.
The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times.
Note that k is guaranteed to be a positive integer.
You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc.
Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k.
For example, there will not be input like 3a or 2[4].
The test cases are generated so that the length of the output will never exceed 105.
*/

/*
{"3[a]2[bc]", "aaabcbc"},
{"3[a2[c]]", "accaccacc"},
{"2[abc]3[cd]ef", "abcabccdcdcdef"},
*/

func decodeString(s string) string {
	var curRes strings.Builder
	var numStack IntStack
	var strStack StrStack
	var idx int
	for idx < len(s) {
		if isDigit(s[idx]) { // Start of multiplier
			var count int
			for isDigit(s[idx]) { // Getting the whole number
				count = count*10 + int(s[idx]-'0')
				idx++
			}
			numStack = append(numStack, count)
		} else if s[idx] == '[' { // Start of multiplier value
			strStack = append(strStack, curRes.String())
			curRes.Reset()
			idx++
		} else if s[idx] == ']' { // End of mulltiplier value
			// Multiply the current res with the multiplier in the stack
			temp := strings.Repeat(curRes.String(), numStack.Pop())
			// Make current res the previous result + what we got
			curRes.Reset()
			curRes.WriteString(strStack.Pop())
			curRes.WriteString(temp)
			idx++
		} else { // Regular letters
			// Just add current letter to current res
			curRes.WriteByte(s[idx])
			idx++
		}
	}
	return curRes.String()
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

type IntStack []int

func (s *IntStack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

type StrStack []string

func (s *StrStack) Pop() string {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func decodeStringRecursive(s string) string {
	var i int
	return recursiveHelper(&s, &i)
}

func recursiveHelper(s *string, i *int) string {
	var res strings.Builder

	for *i < len(*s) && (*s)[*i] != ']' {
		if !isDigit((*s)[*i]) {
			res.WriteByte((*s)[*i])
			*i++
		} else {
			var count int
			for isDigit((*s)[*i]) { // Getting the whole number
				count = count*10 + int((*s)[*i]-'0')
				*i++
			}
			*i++
			temp := recursiveHelper(s, i)
			*i++
			res.WriteString(strings.Repeat(temp, count))
		}
	}
	return res.String()
}
