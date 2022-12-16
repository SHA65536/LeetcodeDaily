package problem1137

/*
The Tribonacci sequence Tn is defined as follows:
T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
Given n, return the value of Tn.
*/

func tribonacci(n int) int {
	var a, b, c int = 0, 1, 1
	if n == 0 {
		return 0
	} else if n < 3 {
		return 1
	}
	for ; n > 2; n-- {
		t := a + b + c
		a = b
		b = c
		c = t
	}
	return c
}
