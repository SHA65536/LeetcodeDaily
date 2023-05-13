package problem2466

/*
Given the integers zero, one, low, and high, we can construct a string by starting with an empty string,
and then at each step perform either of the following:
    Append the character '0' zero times.
    Append the character '1' one times.
This can be performed any number of times.
A good string is a string constructed by the above process having a length between low and high (inclusive).
Return the number of different good strings that can be constructed satisfying these properties.
Since the answer can be large, return it modulo 10^9 + 7.
*/

const mod int = 1e9 + 7

func countGoodStrings(low, high, zero, one int) int {
	var res int
	// dp[i] represents the number of strings that can be formed
	// with length i
	var dp = make([]int, high+1)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		// If we could've put zeros on the last step
		if i >= zero {
			// Add the last step's cached to the current cached
			dp[i] = (dp[i] + dp[i-zero]) % mod
		}
		// If we could've put ones on the last step
		if i >= one {
			// Add the last step's cached to the current cached
			dp[i] = (dp[i] + dp[i-one]) % mod
		}
		// If the length is between low and high
		if i >= low {
			// Update the result with the current numbers
			res = (res + dp[i]) % mod
		}
	}
	return res
}
