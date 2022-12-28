package problem1962

/*
You are given a 0-indexed integer array piles, where piles[i] represents the number of stones in the ith pile, and an integer k.
You should apply the following operation exactly k times:
Choose any piles[i] and remove floor(piles[i] / 2) stones from it.
Notice that you can apply the operation on the same pile more than once.
Return the minimum possible total number of stones remaining after applying the k operations.
floor(x) is the greatest integer that is smaller than or equal to x (i.e., rounds x down).
*/

import "container/heap"

func minStoneSum(piles []int, k int) int {
	var res int
	var pileHeap = MakeHeap(MaxHeap)
	for i := range piles {
		heap.Push(pileHeap, piles[i])
	}
	for ; k > 0; k-- {
		v := heap.Pop(pileHeap).(int)
		v -= v / 2
		heap.Push(pileHeap, v)
	}
	for pileHeap.Len() != 0 {
		res += heap.Pop(pileHeap).(int)
	}
	return res
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
