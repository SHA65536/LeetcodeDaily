package problem1323

import (
	"math"
	"strconv"
)

/*
You are given a positive integer num consisting only of digits 6 and 9.
Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).
*/

func maximum69Number(num int) int {
	var count int
	str := strconv.Itoa(num)
	for i := range str {
		if str[i] == '6' {
			count = len(str) - i
			break
		}
	}
	return num + (3 * int(math.Pow10(count-1)))
}

func maximum69NumberOpt(num int) int {
	switch {
	case num/1000 == 6:
		return num + 3000
	case (num/100)%10 == 6:
		return num + 300
	case (num/10)%10 == 6:
		return num + 30
	case num%10 == 6:
		return num + 3
	}
	return num
}
