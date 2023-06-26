package problem2462

import (
	"container/heap"
)

/*
You are given a 0-indexed integer array costs where costs[i] is the cost of hiring the ith worker.
You are also given two integers k and candidates. We want to hire exactly k workers according to the following rules:
    You will run k sessions and hire exactly one worker in each session.
    In each hiring session, choose the worker with the lowest cost from either the first candidates workers or the last candidates workers.
	Break the tie by the smallest index.
        For example, if costs = [3,2,7,7,1,2] and candidates = 2, then in the first hiring session,
		we will choose the 4th worker because they have the lowest cost [3,2,7,7,1,2].
        In the second hiring session, we will choose 1st worker because they have the same lowest cost as
		4th worker but they have the smallest index [3,2,7,7,2].
		Please note that the indexing may be changed in the process.
    If there are fewer than candidates workers remaining, choose the worker with the lowest cost among them. Break the tie by the smallest index.
    A worker can only be chosen once.
Return the total cost to hire exactly k workers.
*/

func totalCost(costs []int, k, candidates int) int64 {
	var res int64

	if len(costs) == 1 {
		return int64(costs[0])
	}

	// left and right represent the inner index of the left and right groups
	left, right := 0, len(costs)-1

	// lPq and rPq are priority que
	lPq, rPq := MakeHeap(MinHeap), MakeHeap(MinHeap)
	for i := 0; i < candidates && right > left; i++ {
		heap.Push(lPq, costs[i])
		heap.Push(rPq, costs[len(costs)-i-1])
		left++
		right--
	}

	for ; k > 0; k-- {
		if lPq.Len() == 0 {
			// If left is empty, take from right
			res += int64(heap.Pop(rPq).(int))
			continue
		} else if rPq.Len() == 0 {
			// If right is empty, take from left
			res += int64(heap.Pop(lPq).(int))
			continue
		}

		lTemp, rTemp := lPq.Peek(), rPq.Peek()
		if rTemp < lTemp {
			// Take from right
			res += int64(heap.Pop(rPq).(int))
			if right >= left {
				// Only push if the groups aren't colliding
				heap.Push(rPq, costs[right])
				right--
			}
		} else {
			// Take from left
			res += int64(heap.Pop(lPq).(int))
			if right >= left {
				// Only push if the groups aren't colliding
				heap.Push(lPq, costs[left])
				left++
			}
		}
	}

	return res
}

// Priority queue implementation
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

func MaxHeap(i, j int) bool { return i > j }
func MinHeap(i, j int) bool { return i < j }
