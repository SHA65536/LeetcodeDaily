package problem2558

/*
You are given an integer array gifts denoting the number of gifts in various piles. Every second, you do the following:
    Choose the pile with the maximum number of gifts.
    If there is more than one pile with the maximum number of gifts, choose any.
    Leave behind the floor of the square root of the number of gifts in the pile. Take the rest of the gifts.
Return the number of gifts remaining after k seconds.
*/

import (
	"container/heap"
	"math"
)

func pickGifts(gifts []int, k int) int64 {
	var res int64
	var giftQueue = MakeHeap(MaxHeap)
	for i := range gifts {
		heap.Push(giftQueue, gifts[i])
	}
	for ; k > 0; k-- {
		val := heap.Pop(giftQueue).(int)
		heap.Push(giftQueue, int(math.Sqrt(float64(val))))
	}
	for giftQueue.Len() > 0 {
		val := heap.Pop(giftQueue).(int)
		res += int64(val)
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
