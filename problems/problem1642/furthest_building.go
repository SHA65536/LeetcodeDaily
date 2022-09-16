package problem1642

/*
https://leetcode.com/problems/furthest-building-you-can-reach

You are given an integer array heights representing the heights of buildings, some bricks, and some ladders.
You start your journey from building 0 and move to the next building by possibly using bricks or ladders.
While moving from building i to building i+1 (0-indexed),
If the current building's height is greater than or equal to the next building's height, you do not need a ladder or bricks.
If the current building's height is less than the next building's height, you can either use one ladder or (h[i+1] - h[i]) bricks.
Return the furthest building index (0-indexed) you can reach if you use the given ladders and bricks optimally.
*/

import (
	"container/heap"
)

func furthestBuilding(heights []int, bricks int, ladders int) int {
	// diffs is out priority queue for smallest differences
	diffs := &IntHeap{}
	heap.Init(diffs)
	for i := 1; i < len(heights); i++ {
		// ignoring buildings that we don't need to climb
		if heights[i-1] >= heights[i] {
			continue
		}
		// adding current difference to the queue
		heap.Push(diffs, heights[i]-heights[i-1])
		// if we're out of ladders, we replace the shortest ladder with bricks
		if diffs.Len() > ladders {
			smallest := heap.Pop(diffs).(int)
			// if we don't have bricks to replace the shortest, we're done
			if smallest > bricks {
				return i - 1
			}
			bricks -= smallest
		}
	}
	return len(heights) - 1
}

// Priority Queue Implementation using container/heap
// I did not write thise
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
