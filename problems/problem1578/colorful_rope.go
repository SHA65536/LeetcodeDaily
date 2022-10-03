package problem1578

/*
Alice has n balloons arranged on a rope. You are given a 0-indexed string colors where colors[i] is the color of the ith balloon.
Alice wants the rope to be colorful.
She does not want two consecutive balloons to be of the same color, so she asks Bob for help.
Bob can remove some balloons from the rope to make it colorful.
You are given a 0-indexed integer array neededTime where neededTime[i] is the time (in seconds) that Bob needs to remove the ith balloon from the rope.
*/

func minCost(colors string, neededTime []int) int {
	var prev, cur, res int
	cur = 1
	for cur < len(colors) {
		if colors[cur] == colors[prev] {
			if neededTime[cur] < neededTime[prev] {
				res += neededTime[cur]
				cur++
			} else {
				res += neededTime[prev]
				prev = cur
				cur++
			}
		} else {
			prev = cur
			cur++
		}
	}
	return res
}
