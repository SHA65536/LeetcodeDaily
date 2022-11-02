package problem0709

import "strings"

/*
Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.
*/

func toLowerCase(s string) string {
	builder := strings.Builder{}
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			builder.WriteByte(s[i] - 'A' + 'a')
		} else {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}
