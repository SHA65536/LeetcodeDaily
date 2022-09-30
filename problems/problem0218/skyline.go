package problem0218

import (
	"container/heap"
	"sort"
)

/*
A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance.
Given the locations and heights of all the buildings, return the skyline formed by these buildings collectively.

The geometric information of each building is given in the array buildings where buildings[i] = [lefti, righti, heighti]:

lefti is the x coordinate of the left edge of the ith building.
righti is the x coordinate of the right edge of the ith building.
heighti is the height of the ith building.
You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.

The skyline should be represented as a list of "key points" sorted by their x-coordinate in the form [[x1,y1],[x2,y2],...].
Each key point is the left endpoint of some horizontal segment in the skyline except the last point in the list,
which always has a y-coordinate 0 and is used to mark the skyline's termination where the rightmost building ends.
Any ground between the leftmost and rightmost buildings should be part of the skyline's contour.
*/

func getSkyline(buildings [][]int) [][]int {
	var res = [][]int{}
	var heights = make([][]int, len(buildings)*2)
	var pq = &MaxHeap{}
	var prevMax, curMax int
	for i := range buildings {
		heights[i*2] = []int{buildings[i][0], -buildings[i][2]}
		heights[(i*2)+1] = []int{buildings[i][1], buildings[i][2]}
	}
	sort.Slice(heights, func(i, j int) bool {
		if heights[i][0] == heights[j][0] {
			return heights[i][1] < heights[j][1]
		}
		return heights[i][0] < heights[j][0]
	})
	heap.Init(pq)
	heap.Push(pq, 0)
	for i := range heights {
		if heights[i][1] < 0 {
			heap.Push(pq, -heights[i][1])
		} else {
			heap.Remove(pq, pq.Index(heights[i][1]))
		}
		curMax = (*pq)[0]
		if curMax != prevMax {
			res = append(res, []int{heights[i][0], curMax})
			prevMax = curMax
		}
	}

	return res
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxHeap) Index(e int) int {
	for i, v := range h {
		if v == e {
			return i
		}
	}
	return -1
}
