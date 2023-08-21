package problem0459

/*
Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.
*/

func repeatedSubstringPattern(s string) bool {
length:
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i != 0 {
			continue
		}
		for j := range s {
			if s[j] != s[j%i] {
				continue length
			}
		}
		return true
	}
	return false
}
