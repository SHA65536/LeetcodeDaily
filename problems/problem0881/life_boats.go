package problem0881

import "sort"

/*
You are given an array people where people[i] is the weight of the ith person,
and an infinite number of boats where each boat can carry a maximum weight of limit.
Each boat carries at most two people at the same time, provided the sum of the weight of those people is at most limit.
Return the minimum number of boats to carry every given person.
*/

func numRescueBoats(people []int, limit int) int {
	var light, heavy, res int
	// Sort from lightest to heaviest
	sort.Ints(people)
	// Start from both ends and approach the middle
	for heavy = len(people) - 1; light <= heavy; res++ {
		if people[light]+people[heavy] <= limit {
			// If both the current heaviest an current lightest can
			// go in the same boat, pair them up
			light++
			heavy--
		} else {
			// If not, heaviest will go alone
			heavy--
		}
	}
	return res
}
