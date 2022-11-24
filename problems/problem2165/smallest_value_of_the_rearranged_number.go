package problem2165

import (
	"sort"
	"strconv"
)

/*
You are given an integer num. Rearrange the digits of num such that its value is minimized and it does not contain any leading zeros.
Return the rearranged number with minimal value.
Note that the sign of the number does not change after rearranging the digits.
*/

func smallestNumber(num int64) int64 {
	res := []byte(strconv.FormatInt(num, 10))
	if num < 0 {
		res[0] = '9'
		sort.Slice(res, func(i, j int) bool {
			return res[i] > res[j]
		})
		res[0] = '-'
	} else {
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		for i := range res {
			if res[i] != '0' {
				res[0], res[i] = res[i], res[0]
				break
			}
		}
	}
	num, _ = strconv.ParseInt(string(res), 10, 64)
	return num
}
