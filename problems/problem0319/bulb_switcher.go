package problem0319

import "math"

/*
There are n bulbs that are initially off. You first turn on all the bulbs, then you turn off every second bulb.
On the third round, you toggle every third bulb (turning on if it's off or turning off if it's on).
For the ith round, you toggle every i bulb. For the nth round, you only toggle the last bulb.
Return the number of bulbs that are on after n rounds.
*/

// For this solution I just ran the problem on bruteforce
// observed the pattern of lit bulbs and mimicked it here
func bulbSwitch(n int) int {
	var res, cur int = 0, 3
	for i := 0; i < n; {
		res++
		i += cur
		cur += 2
	}
	return res
}

func bulbSwitchSqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}
