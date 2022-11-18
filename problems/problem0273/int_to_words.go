package problem0273

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Convert a non-negative integer num to its English words representation.
*/

func numberToWords(num int) string {
	var digits string
	var sb strings.Builder
	if num == 0 {
		return "Zero"
	}
	digits = fmt.Sprintf("%010d", num)
	bils, _ := strconv.Atoi(digits[0:1])
	mils, _ := strconv.Atoi(digits[1:4])
	thus, _ := strconv.Atoi(digits[4:7])
	ones, _ := strconv.Atoi(digits[7:10])
	if bils > 0 {
		sb.WriteString(GetQuantifier(bils))
		sb.WriteString("Billion ")
	}
	if mils > 0 {
		sb.WriteString(GetQuantifier(mils))
		sb.WriteString("Million ")
	}
	if thus > 0 {
		sb.WriteString(GetQuantifier(thus))
		sb.WriteString("Thousand ")
	}
	if ones > 0 {
		sb.WriteString(GetQuantifier(ones))
	}
	digits = sb.String()
	return digits[:len(digits)-1]
}

func GetQuantifier(input int) string {
	var sb strings.Builder
	if input >= 100 {
		sb.WriteString(Words[input/100])
		sb.WriteString(" Hundred ")
	}
	input %= 100
	if input <= 0 {
		return sb.String()
	}
	if input > 19 {
		sb.WriteString(Words[(input/10)*10])
		sb.WriteString(" ")
		input %= 10
		if input > 0 {
			sb.WriteString(Words[input])
			sb.WriteString(" ")
		}
	} else {
		sb.WriteString(Words[input])
		sb.WriteString(" ")
	}
	return sb.String()
}

var Words = map[int]string{
	1: "One", 2: "Two", 3: "Three",
	4: "Four", 5: "Five", 6: "Six",
	7: "Seven", 8: "Eight", 9: "Nine",
	10: "Ten", 11: "Eleven", 12: "Twelve",
	13: "Thirteen", 14: "Fourteen", 15: "Fifteen",
	16: "Sixteen", 17: "Seventeen", 18: "Eighteen",
	19: "Nineteen", 20: "Twenty", 30: "Thirty",
	40: "Forty", 50: "Fifty", 60: "Sixty",
	70: "Seventy", 80: "Eighty", 90: "Ninety",
}

// Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven
// 2|147|483|6|47
