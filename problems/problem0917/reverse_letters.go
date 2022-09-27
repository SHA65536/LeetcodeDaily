package problem0917

/*
Given a string s, reverse the string according to the following rules:
All the characters that are not English letters remain in the same position.
All the English letters (lowercase or uppercase) should be reversed.
Return s after reversing it.
*/

func reverseOnlyLetters(s string) string {
	var res = make([]byte, len(s))
	var end = len(res) - 1
	for i := range s {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')) {
			res[i] = s[i]
		} else {
			for !((s[end] >= 'a' && s[end] <= 'z') || (s[end] >= 'A' && s[end] <= 'Z')) {
				end--
			}
			res[i] = s[end]
			end--
		}
	}
	return string(res)
}
