package problem0326

/*
Given an integer n, return true if it is a power of three. Otherwise, return false.
An integer n is a power of three, if there exists an integer x such that n == 3x.
*/

const MaxPowreOfThree int = 1162261467

func isPowerOfThree(n int) bool {
	return n > 0 && MaxPowreOfThree%n == 0
}
