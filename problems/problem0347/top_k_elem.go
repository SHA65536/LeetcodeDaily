package problem0347

import (
	"container/heap"
	"sort"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in any order.
*/

func topKFrequentSort(nums []int, k int) []int {
	var freq = map[int]int{}
	var res []int

	// Calculate frequencies
	for i := range nums {
		if _, ok := freq[nums[i]]; !ok {
			res = append(res, nums[i])
		}
		freq[nums[i]]++
	}

	// Sort by frequency descending
	sort.Slice(res, func(i, j int) bool {
		return freq[nums[i]] > freq[nums[j]]
	})

	// Get 5 biggest items
	return res[:k]
}

func topKFrequentHeap(nums []int, k int) []int {
	var freq = map[int]int{}
	var res []int

	// Calculate frequencies
	for i := range nums {
		freq[nums[i]]++
	}

	h := &FreqHeap{}
	var i int
	// Add frequencies to heap
	for key, val := range freq {
		if i < k {
			// Add initial k items
			heap.Push(h, [2]int{val, key})
		} else if val > (*h)[0][0] {
			// If current item bigger than the smallest
			// item in the heap, pop the smallest, add the new one
			heap.Pop(h)
			heap.Push(h, [2]int{val, key})
		}
		i++
	}

	// Move items from heap to res
	res = make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).([2]int)[1]
	}
	return res
}

// Implementation of a min-heap for integers
type FreqHeap [][2]int

func (h FreqHeap) Len() int {
	return len(h)
}

func (h FreqHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h FreqHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FreqHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *FreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
