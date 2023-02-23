package problem0502

import (
	"container/heap"
	"sort"
)

/*
Suppose LeetCode will start its IPO soon. In order to sell a good price of its shares to Venture Capital,
LeetCode would like to work on some projects to increase its capital before the IPO.
Since it has limited resources, it can only finish at most k distinct projects before the IPO.
Help LeetCode design the best way to maximize its total capital after finishing at most k distinct projects.
You are given n projects where the ith project has a pure profit profits[i] and a minimum capital of capital[i] is needed to start it.
Initially, you have w capital. When you finish a project, you will obtain its pure profit and the profit will be added to your total capital.
Pick a list of at most k distinct projects from given projects to maximize your final capital, and return the final maximized capital.
The answer is guaranteed to fit in a 32-bit signed integer.
*/

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	var order = make([]int, len(profits))
	for i := range order {
		order[i] = i
	}

	// Sorting projects by increasing capital
	sort.Slice(order, func(i, j int) bool {
		return capital[order[i]] < capital[order[j]]
	})

	// pq Holds the projects we can start, from most profit to least
	var pq = MakeHeap(MaxHeap)

	// Looping each day
	var i int
	for ; k > 0; k-- {
		// Adding the projects we can start into the pq
		for i < len(order) && capital[order[i]] <= w {
			heap.Push(pq, profits[order[i]])
			i++
		}
		// Each day we take the most profitable project
		if pq.Len() > 0 {
			r := heap.Pop(pq).(int)
			w += r
		} else {
			break
		}
	}

	return w
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

func MaxHeap(i, j int) bool { return i > j }
func MinHeap(i, j int) bool { return i < j }
