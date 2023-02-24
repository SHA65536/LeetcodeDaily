package problem1675

import "container/heap"

/*
You are given an array nums of n positive integers.
You can perform two types of operations on any element of the array any number of times:
    If the element is even, divide it by 2.
        For example, if the array is [1,2,3,4], then you can do this operation on the last element, and the array will be [1,2,3,2].
    If the element is odd, multiply it by 2.
        For example, if the array is [1,2,3,4], then you can do this operation on the first element, and the array will be [2,2,3,4].
The deviation of the array is the maximum difference between any two elements in the array.
Return the minimum deviation the array can have after performing some number of operations.
*/

const MAX_N = 1e9

func minimumDeviation(nums []int) int {
	var res, small int = MAX_N, MAX_N
	var temp int
	var pq = MakeHeap(MaxHeap)

	// Making all odd numbers even and finding the minimum
	for i := range nums {
		temp = nums[i]
		if nums[i]%2 == 1 {
			temp *= 2
		}
		heap.Push(pq, temp)
		small = min(temp, small)
	}

	// Trying to divide the largest number until you can't
	for pq.Peek()%2 == 0 {
		res = min(res, pq.Peek()-small)
		small = min(small, pq.Peek()/2)
		heap.Push(pq, pq.Peek()/2)
		heap.Pop(pq)
	}

	return min(res, pq.Peek()-small)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
