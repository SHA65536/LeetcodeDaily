package problem0664

/*
There is a strange printer with the following two special properties:
The printer can only print a sequence of the same character each time.
At each turn, the printer can print new characters starting from and ending at any place and will cover the original existing characters.
Given a string s, return the minimum number of turns the printer needed to print it.
*/

func strangePrinter(s string) int {
	// Remove duplicate characters to reduce N
	var lst = make([]byte, 1, len(s))
	lst[0] = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			lst = append(lst, s[i])
		}
	}

	// dp[i][j] is number of turns we need to print s[i:j]
	var dp = make([][]int, len(lst))
	for i := range dp {
		dp[i] = make([]int, len(lst))
		// A single character takes one turn
		dp[i][i] = 1
		// Two characters take two turns
		if i < len(lst)-1 {
			dp[i][i+1] = 2
		}
	}

	// Iterate over every possible length of string
	for l := 2; l < len(lst); l++ {
		// Starting at start
		for st := 0; st+l < len(lst); st++ {
			// The most costly way to print is one turn for each character
			dp[st][st+l] = l + 1
			// Looping over possible ways to divide the string
			for k := 0; k < l; k++ {
				// Take the sum of the two ways
				t := dp[st][st+k] + dp[st+k+1][st+l]
				if lst[st+k] == lst[st+l] {
					// If the edges are the same, we can save a turn
					dp[st][st+l] = min(dp[st][st+l], t-1)
				} else {
					dp[st][st+l] = min(dp[st][st+l], t)
				}

			}
		}
	}

	return dp[0][len(lst)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
