package problem0038

import (
	"strconv"
	"strings"
)

/*
The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
countAndSay(1) = "1"
countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.
To determine how you "say" a digit string, split it into the minimal number of substrings such that each substring contains exactly one unique digit.
Then for each substring, say the number of digits, then say the digit. Finally, concatenate every said digit.
*/

func countAndSay(n int) string {
	var previous string
	var res strings.Builder
	if n == 1 {
		return "1"
	}
	// Getting the output of last number
	previous = countAndSay(n - 1)

	// Numbers represents the number of times a char
	// has been seen in a row
	var numbers = make([]int, 1, 64)
	numbers[0] = 1

	// Bytes represents the corresponding number for
	// the numbers slice
	var bytes = make([]byte, 1, 64)
	bytes[0] = previous[0]

	var cur int
	for i := 1; i < len(previous); i++ {
		if previous[i] == previous[i-1] {
			// If the number is the same as the last one
			// increase the count
			numbers[cur]++
		} else {
			// If it's a new number, create a new count for it
			numbers = append(numbers, 1)
			bytes = append(bytes, previous[i])
			cur++
		}
	}
	// Converting both arrays to string
	for i := range numbers {
		res.WriteString(strconv.Itoa(numbers[i]))
		res.WriteByte(bytes[i])
	}
	return res.String()
}
