package problem1354

import (
	"container/heap"
)

/*
https://leetcode.com/problems/construct-target-array-with-multiple-sums

You are given an array target of n integers. From a starting array arr consisting of n 1's, you may perform the following procedure:
let x be the sum of all elements currently in your array.
choose index i, such that 0 <= i < n and set the value of arr at index i to x.
You may repeat this procedure as many times as needed.
Return true if it is possible to construct the target array from arr, otherwise, return false.
*/

func isPossible(target []int) bool {
	var pq = &MaxHeap{}
	var totalSum, num int
	heap.Init(pq)
	for _, num := range target {
		heap.Push(pq, num)
		totalSum += num
	}
	for num = heap.Pop(pq).(int); true; num = heap.Pop(pq).(int) {
		totalSum -= num
		if num == 1 || totalSum == 1 {
			return true
		}
		if num < totalSum || totalSum == 0 || num%totalSum == 0 {
			return false
		}
		num = num % totalSum
		totalSum += num
		heap.Push(pq, num)
	}
	return false
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
