package problem1710

import (
	"sort"
)

/*
https://leetcode.com/problems/maximum-units-on-a-truck/

You are assigned to put some amount of boxes onto one truck. You are given a 2D array boxTypes, where boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi]:
numberOfBoxesi is the number of boxes of type i.
numberOfUnitsPerBoxi is the number of units in each box of the type i.
You are also given an integer truckSize, which is the maximum number of boxes that can be put on the truck.
You can choose any boxes to put on the truck as long as the number of boxes does not exceed truckSize.
Return the maximum total number of units that can be put on the truck.
*/

func maximumUnits(boxTypes [][]int, truckSize int) int {
	var total int
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for i := 0; i < len(boxTypes) && truckSize > 0; i++ {
		if truckSize > boxTypes[i][0] {
			total += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			total += boxTypes[i][1] * truckSize
			break
		}
	}
	return total
}
