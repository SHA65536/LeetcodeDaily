package problem0202

/*
Write an algorithm to determine if a number n is happy.
A happy number is a number defined by the following process:
    Starting with any positive integer, replace the number by the sum of the squares of its digits.
    Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
    Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.
*/

func isHappy(n int) bool {
	var visited = map[int]bool{n: true}
	for n != 1 {
		n = nextNum(n)
		if visited[n] {
			return false
		}
		visited[n] = true
	}
	return true
}

func nextNum(n int) int {
	var res int
	for n > 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}
