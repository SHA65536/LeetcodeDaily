package problem0899

import "sort"

/*
You are given a string s and an integer k. You can choose one of the first k letters of s and append it at the end of the string.
Return the lexicographically smallest string you could have after applying the mentioned step any number of moves.
*/

func orderlyQueue(s string, k int) string {
	var sliced = []byte(s)
	var minStart int
	if k > 1 { // If K is bigger than 1, you can always reorder it however you want
		sort.Slice(sliced, func(i, j int) bool { return sliced[i] < sliced[j] })
		return string(sliced)
	}
	for i := 1; i < len(sliced); i++ {
		// We compare the string starting at the current min
		// with the string starting at i
		if compareStringByStartingPoints(sliced, minStart, i) {
			minStart = i
		}
	}
	// Return the string starting at minStart
	return string(sliced[minStart:]) + string(sliced[:minStart])
}

func compareStringByStartingPoints(input []byte, s1, s2 int) bool {
	var l = len(input)
	for i := range input {
		if input[(s1+i)%l] > input[(s2+i)%l] {
			return true
		} else if input[(s1+i)%l] < input[(s2+i)%l] {
			return false
		}
	}
	return false
}
