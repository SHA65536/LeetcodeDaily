package problem1338

import (
	"sort"
)

/*
You are given an integer array arr.
You can choose a set of integers and remove all the occurrences of these integers in the array.
Return the minimum size of the set so that at least half of the integers of the array are removed.
*/

func minSetSize(arr []int) int {
	var half, sum int
	var numsMap = map[int]int{}
	var numsFreq = []int{}
	for _, num := range arr {
		if _, ok := numsMap[num]; !ok {
			numsFreq = append(numsFreq, num)
		}
		numsMap[num]++
	}
	sort.Slice(numsFreq, func(i, j int) bool {
		return numsMap[numsFreq[i]] > numsMap[numsFreq[j]]
	})
	half = len(arr) / 2
	for i := range numsFreq {
		sum += numsMap[numsFreq[i]]
		if sum >= half {
			return i + 1
		}
	}
	return len(arr)
}
