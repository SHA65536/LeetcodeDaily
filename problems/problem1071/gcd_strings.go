package problem1071

/*
For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).
Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.
*/

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}
	for i := len(str1); i > 0; i-- {
		div := str1[:i]
		if len(str1)%len(div) == 0 && len(str2)%len(div) == 0 {
			if canDivide(str1, div) && canDivide(str2, div) {
				return div
			}
		}
	}
	return ""
}

func canDivide(str, div string) bool {
	for i := range str {
		if str[i] != div[i%len(div)] {
			return false
		}
	}
	return true
}
