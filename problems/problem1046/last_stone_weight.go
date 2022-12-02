package problem1046

import (
	"container/heap"
)

/*
You are given an array of integers stones where stones[i] is the weight of the ith stone.
We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:
If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.
Return the weight of the last remaining stone. If there are no stones left, return 0.
*/

func lastStoneWeight(stones []int) int {
	var stoneHeap = MakeHeap(func(i1, i2 int) bool { return i1 > i2 })
	// Loading stones into heap
	for i := range stones {
		heap.Push(stoneHeap, stones[i])
	}
	// Smashing untill 1 left
	for stoneHeap.Len() > 1 {
		x := heap.Pop(stoneHeap).(int)
		y := heap.Pop(stoneHeap).(int)
		if x > y {
			heap.Push(stoneHeap, x-y)
		} else if y > x {
			heap.Push(stoneHeap, y-x)
		}
	}
	if stoneHeap.Len() == 0 {
		return 0
	}
	return heap.Pop(stoneHeap).(int)
}

type Heap struct {
	Values   []int
	LessFunc func(int, int) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() int          { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(int)) }

func MakeHeap(less func(int, int) bool) *Heap {
	return &Heap{LessFunc: less}
}
