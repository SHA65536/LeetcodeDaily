package problem1523

/*
Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).
*/

func countOdds(low int, high int) int {
	switch (low % 2) + (high % 2) {
	case 2: // Both odd
		return (high-low)/2 + 1
	case 1: // One odd
		return (high-low)/2 + 1
	case 0: // Both Even
		return (high - low) / 2
	}
	return -1
}
