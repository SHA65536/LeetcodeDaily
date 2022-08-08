package problem1220

/*
Given an integer n, your task is to count how many strings of length n can be formed under the following rules:
Each character is a lower case vowel ('a', 'e', 'i', 'o', 'u')
Each vowel 'a' may only be followed by an 'e'.
Each vowel 'e' may only be followed by an 'a' or an 'i'.
Each vowel 'i' may not be followed by another 'i'.
Each vowel 'o' may only be followed by an 'i' or a 'u'.
Each vowel 'u' may only be followed by an 'a'.
Since the answer may be too large, return it modulo 10^9 + 7.
*/

const max = 1000000007
const (
	A = iota
	E
	I
	O
	U
)

var rules = [5][]int{
	{E}, {A, I}, {A, E, O, U}, {I, U}, {A},
}

func countVowelPermutationNaive(n int) int {
	var res int
	for i := A; i <= U; i++ {
		res += countVowelsNaive(i, n-1)
	}
	return res % max
}

func countVowelsNaive(prev int, n int) int {
	var res int
	if n == 0 {
		return 1
	}
	allowed := rules[prev]
	for _, i := range allowed {
		res += countVowelsNaive(i, n-1)
	}
	return res
}

func countVowelPermutation(n int) int {
	// actual sums
	var a, e, i, o, u int = 1, 1, 1, 1, 1
	// temporary sums
	var at, et, it, ot, ut int
	for j := 1; j < n; j++ {
		at, et, it, ot, ut = a, e, i, o, u
		a = (et + it + ut) % max
		e = (at + it) % max
		i = (et + ot) % max
		o = it % max
		u = (it + ot) % max
	}
	return (a + e + i + o + u) % max
}
