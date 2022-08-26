package problem0869

import (
	"math"
	"strconv"
)

/*
You are given an integer n.
We reorder the digits in any order (including the original order) such that the leading digit is not zero.
Return true if and only if we can do this so that the resulting number is a power of two.
*/

func reorderedPowerOf2(n int) bool {
	var numFreq [10]int
	var numStr = strconv.Itoa(n)
	var highLimit = int(math.Pow10(len(numStr)))
	var lowLimit = highLimit / 10
	for i := range numStr {
		numFreq[numStr[i]-'0']++
	}
	for i := 1; i < highLimit; i <<= 1 {
		if i >= lowLimit {
			var curStr = strconv.Itoa(i)
			var curFreq [10]int
			for j := range curStr {
				curFreq[curStr[j]-'0']++
			}
			if curFreq == numFreq {
				return true
			}
		}
	}
	return false
}
