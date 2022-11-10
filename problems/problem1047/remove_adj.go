package problem1047

/*
You are given a string s consisting of lowercase English letters.
A duplicate removal consists of choosing two adjacent and equal letters and removing them.
We repeatedly make duplicate removals on s until we no longer can.
Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.
*/

func removeDuplicates(s string) string {
	var res = []byte(s)
	var cur int
	for i := 0; i < len(res); i++ {
		if cur > 0 && res[i] == res[cur-1] {
			cur--
		} else {
			res[cur] = res[i]
			cur++
		}
	}
	return string(res[0:cur])
}
