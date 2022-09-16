package problem2007

import "sort"

/*
An integer array original is transformed into a doubled array changed by appending twice the value of every element in original,
and then randomly shuffling the resulting array.
Given an array changed, return original if changed is a doubled array.
If changed is not a doubled array, return an empty array.
The elements in original may be returned in any order.
*/

func findOriginalArray(changed []int) []int {
	var res = []int{}
	var needed = map[int]int{}
	// Checking for even number of numbers
	if len(changed)%2 != 0 {
		return []int{}
	}
	// Sorting in ascending order
	sort.Ints(changed)
	for i := range changed {
		if val := needed[changed[i]]; val > 0 {
			// If value is needed as a double, continue
			needed[changed[i]]--
		} else {
			// If value is not needed as a double, add to res and request it's double
			res = append(res, changed[i])
			needed[changed[i]*2]++
		}
	}
	// Making sure all needs are satisfied
	for i := range needed {
		if needed[i] != 0 {
			return []int{}
		}
	}
	return res
}
