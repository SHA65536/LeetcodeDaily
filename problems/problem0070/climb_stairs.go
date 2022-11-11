package problem0070

/*
You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps.
In how many distinct ways can you climb to the top?
*/

func climbStairs(n int) int {
	var prev, cur int
	prev, cur = 1, 1
	for n > 1 {
		prev, cur = cur, cur+prev
		n--
	}
	return cur
}
