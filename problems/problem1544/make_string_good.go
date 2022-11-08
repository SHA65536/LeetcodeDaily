package problem1544

/*
Given a string s of lower and upper case English letters.
A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:
0 <= i <= s.length - 2
s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice-versa.
To make the string good, you can choose two adjacent characters that make the string bad and remove them.
You can keep doing this until the string becomes good.
Return the string after making it good. The answer is guaranteed to be unique under the given constraints.
Notice that an empty string is also good.
*/

const offset int = 32

func makeGood(s string) string {
	var res = []byte(s)
	var cur int
	for i := 0; i < len(res); i++ {
		if cur > 0 && absDif(res[i], res[cur-1]) == offset {
			cur--
		} else {
			res[cur] = res[i]
			cur++
		}
	}
	return string(res[0:cur])
}

func absDif(a, b byte) int {
	c := int(a) - int(b)
	if c < 0 {
		return -c
	}
	return c
}
