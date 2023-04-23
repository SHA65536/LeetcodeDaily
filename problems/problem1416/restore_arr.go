package problem1416

/*
A program was supposed to print an array of integers.
The program forgot to print whitespaces and the array is printed as a string of digits s and all we know is that all integers
in the array were in the range [1, k] and there are no leading zeros in the array.
Given the string s and the integer k, return the number of the possible arrays that can be printed as s using the mentioned program.
Since the answer may be very large, return it modulo 109 + 7.
*/

const MOD int = 1e9 + 7

func numberOfArrays(s string, limit int) int {
	// cache[i] stores the number of arrays possible from
	// the string s[i:]
	var cache = make([]int, len(s))
	// calcArrays calculates the number of arrays possible from
	// the string s[i:]
	var calcArrays func(idx int) int

	calcArrays = func(idx int) int {
		var res, i, n int
		// If we're overshooting or the number starts with 0
		if idx == len(s) || s[idx] == '0' {
			return 0
		}
		// If we already calculated this string
		if cache[idx] != 0 {
			return cache[idx]
		}
		// Loop until end of string or n > limit
		for i = idx; i < len(s); i++ {
			// Make a number from the start of the string
			n = (n * 10) + int(s[i]-'0')
			// Until the number is invalid (bigger than limit)
			if n > limit {
				break
			}
			// Add the rest of the string to the result
			res += calcArrays(i+1) % MOD
		}

		// If we reached the end of the string, add 1
		if i == len(s) && n <= limit {
			res++
		}

		// Save to cache
		cache[idx] = res % MOD
		return cache[idx]
	}

	return calcArrays(0)
}
