package problem0093

import (
	"strconv"
	"strings"
)

/*
A valid IP address consists of exactly four integers separated by single dots.
Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.
For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s.
You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.
*/

func restoreIpAddresses(s string) []string {
	if len(s) > 12 || len(s) <= 0 {
		return []string{}
	}
	return restoreHelper(s, 4)
}

func restoreHelper(s string, cnt int) []string {
	var res = []string{}
	if len(s) == 0 {
		return res
	}
	if cnt == 1 {
		if validNum(s) {
			return []string{s}
		}
		return res
	}
	for i := 1; i < 4 && i < len(s); i++ {
		if validNum(s[0:i]) {
			temp := restoreHelper(s[i:], cnt-1)
			for _, t := range temp {
				res = append(res, s[0:i]+"."+t)
			}
		}
	}
	return res
}

func validNum(s string) bool {
	if len(s) == 0 || len(s) > 3 || (len(s) > 1 && strings.HasPrefix(s, "0")) {
		return false
	}
	if val, _ := strconv.Atoi(s); val > 255 {
		return false
	}
	return true
}
