package problem2390

/*
You are given a string s, which contains stars *.
In one operation, you can:
    Choose a star in s.
    Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.
Note:
    The input will be generated such that the operation is always possible.
    It can be shown that the resulting string will always be unique.
*/

func removeStars(s string) string {
	var res = make([]byte, len(s))
	var rCur, wCur int
	for rCur = 0; rCur < len(s); rCur++ {
		if s[rCur] == '*' {
			wCur--
		} else {
			res[wCur] = s[rCur]
			wCur++
		}
	}
	return string(res[:wCur])
}
