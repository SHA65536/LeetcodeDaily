package problem0848

/*
You are given a string s of lowercase English letters and an integer array shifts of the same length.
Call the shift() of a letter, the next letter in the alphabet, (wrapping around so that 'z' becomes 'a').
	For example, shift('a') = 'b', shift('t') = 'u', and shift('z') = 'a'.
Now for each shifts[i] = x, we want to shift the first i + 1 letters of s, x times.
Return the final string after all such shifts to s are applied.
*/

func shiftingLetters(s string, shifts []int) string {
	var temp = []byte(s)
	var sum int
	for i := len(shifts) - 1; i >= 0; i-- {
		sum = (sum + shifts[i]) % 26
		temp[i] = ((temp[i] - 'a' + byte(sum)) % 26) + 'a'
	}
	return string(temp)
}
