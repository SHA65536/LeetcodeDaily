package problem1465

import "sort"

/*
https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/

You are given a rectangular cake of size h x w and two arrays of integers horizontalCuts and verticalCuts where:
horizontalCuts[i] is the distance from the top of the rectangular cake to the ith horizontal cut and similarly, and
verticalCuts[j] is the distance from the left of the rectangular cake to the jth vertical cut.
Return the maximum area of a piece of cake after you cut at each horizontal and vertical position provided in the arrays horizontalCuts and verticalCuts.
Since the answer can be a large number, return this modulo 10^9 + 7.
*/

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	var prev, maxH, maxV int
	// Sorting the cuts so we can calculate biggest difference between cuts
	sort.Ints(horizontalCuts)
	horizontalCuts = append(horizontalCuts, h)
	sort.Ints(verticalCuts)
	verticalCuts = append(verticalCuts, w)
	// Finding biggest horizontal difference between 2 cuts
	for i := 0; i < len(horizontalCuts); i++ {
		if maxH < horizontalCuts[i]-prev {
			maxH = horizontalCuts[i] - prev
		}
		prev = horizontalCuts[i]
	}
	prev = 0
	// Finding biggest vertical difference between 2 cuts
	for i := 0; i < len(verticalCuts); i++ {
		if maxV < verticalCuts[i]-prev {
			maxV = verticalCuts[i] - prev
		}
		prev = verticalCuts[i]
	}
	// Question wanted to return the size modulo 10^9 + 7 :shrug:
	return (maxH * maxV) % (1000000007)
}
