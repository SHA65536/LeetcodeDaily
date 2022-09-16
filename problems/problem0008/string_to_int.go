package problem0008

import (
	"math"
	"regexp"
)

/*
https://leetcode.com/problems/string-to-integer-atoi/

Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

The algorithm for myAtoi(string s) is as follows:

1. Read in and ignore any leading whitespace.
2. Check if the next character (if not already at the end of the string) is '-' or '+'.
	Read this character in if it is either. This determines if the final result is negative or positive respectively.
	Assume the result is positive if neither is present.
3. Read in next the characters until the next non-digit character or the end of the input is reached.
	The rest of the string is ignored.
4. Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32).
	If no digits were read, then the integer is 0.
	Change the sign as necessary (from step 2).
5. If the integer is out of the 32-bit signed integer range [-2^31, 2^31 - 1], then clamp the integer so that it remains in the range.
	Specifically, integers less than -2^31 should be clamped to -2^31, and integers greater than 2^31 - 1 should be clamped to 2^31 - 1.
6. Return the integer as the final result.

Note:
Only the space character ' ' is considered a whitespace character.
Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.
*/

// The really weird edge cases and constraints frustrated me
// Leading to this regex :)
var atoiRegex = regexp.MustCompile(`^\s*([+-]?\d+)`)

func myAtoi(s string) int {
	var result int
	var mod = 1
	normal := atoiRegex.FindStringSubmatch(s)
	if normal == nil {
		return 0
	}
	if normal[1][0:1] == "-" {
		mod = -1
		normal[1] = normal[1][1:]
	} else if normal[1][0:1] == "+" {
		normal[1] = normal[1][1:]
	}
	for _, char := range normal[1] {
		if checkOverflow(result, int(char), mod) {
			if mod == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		result = result*10 + (int(char) - 48)
	}
	return result * mod
}

func checkOverflow(res, cur, mod int) bool {
	if mod == 1 {
		return res > math.MaxInt32/10 || (math.MaxInt32-(res*10)) < (cur-48)
	} else {
		return res*mod < math.MinInt32/10 || (math.MinInt32+(res*10) > (cur-48)*mod)
	}
}
