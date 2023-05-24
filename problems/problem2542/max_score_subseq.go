package problem2542

import (
	"container/heap"
	"sort"
)

/*
You are given two 0-indexed integer arrays nums1 and nums2 of equal length n and a positive integer k.
You must choose a subsequence of indices from nums1 of length k.
For chosen indices i0, i1, ..., ik - 1, your score is defined as:
    The sum of the selected elements from nums1 multiplied with the minimum of the selected elements from nums2.
    It can defined simply as: (nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1]).
Return the maximum possible score.
A subsequence of indices of an array is a set that can be derived from the set {0, 1, ..., n-1} by deleting some or no elements.
*/

func maxScore(nums1, nums2 []int, k int) int64 {
	var res, sum int64
	var order = make([]int, len(nums1))
	for i := range order {
		order[i] = i
	}

	sort.Slice(order, func(i, j int) bool {
		return nums2[order[i]] >= nums2[order[j]]
	})

	var pq = MakeHeap(MinHeap)

	for i := range order {
		e, s := nums2[order[i]], nums1[order[i]]
		heap.Push(pq, s)
		sum += int64(s)
		if pq.Len() > k {
			sum -= int64(heap.Pop(pq).(int))
		}
		if pq.Len() == k {
			res = max(res, sum*int64(e))
		}
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

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
