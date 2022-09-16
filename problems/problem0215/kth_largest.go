package problem0215

import (
	"container/heap"
	"sort"
)

/*
https://leetcode.com/problems/kth-largest-element-in-an-array

Given an integer array nums and an integer k, return the kth largest element in the array.
Note that it is the kth largest element in the sorted order, not the kth distinct element.
*/

func findKthLargestSort(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest(nums []int, k int) int {
	bigs := &MinHeap{}
	heap.Init(bigs)

	for i := 0; i < len(nums); i++ {
		if i < k {
			heap.Push(bigs, nums[i])
		} else if nums[i] > (*bigs)[0] {
			heap.Pop(bigs)
			heap.Push(bigs, nums[i])
		}
	}

	return (*bigs)[0]
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
