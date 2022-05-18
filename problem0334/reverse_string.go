package problem0334

/*
https://leetcode.com/problems/reverse-string/

Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.
*/

func reverseString(s []byte) {
	var start, end = 0, len(s) - 1
	for end > start {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
