package problem1402

import "sort"

/*
A chef has collected data on the satisfaction level of his n dishes. Chef can cook any dish in 1 unit of time.
Like-time coefficient of a dish is defined as the time taken to cook that dish including previous dishes multiplied
by its satisfaction level i.e. time[i] * satisfaction[i].
Return the maximum sum of like-time coefficient that the chef can obtain after dishes preparation.
Dishes can be prepared in any order and the chef can discard some dishes to get this maximum value.
*/

func maxSatisfaction(sat []int) int {
	sort.Ints(sat)
	var res, total int
	// Looping from most satesfying to least satesfying
	// until the satisfaction from the previous dishes is not positive
	for i := len(sat) - 1; i >= 0 && sat[i]+total > 0; i-- {
		// Each time we add the current dish satisfaction
		total += sat[i]
		// And gain all the satisfaction from the previous dishes
		res += total
	}
	return res
}
