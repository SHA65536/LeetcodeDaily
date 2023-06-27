package problem0373

import (
	"container/heap"
	"sort"
)

/*
You are given two integer arrays nums1 and nums2 sorted in ascending order and an integer k.
Define a pair (u, v) which consists of one element from the first array and one element from the second array.
Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.
*/

type Pair struct {
	A, B int // Elements
	Idx  int // Index of the 2nd element
}

func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	var res [][]int
	// Priority queue containing the possible pairs
	var pq = MakeHeap(MinHeap)
	// Pushing all the pairs with nums2[0] as one of the elements
	for i := 0; i < k && i < len(nums1); i++ {
		heap.Push(pq, Pair{nums1[i], nums2[0], 0})
	}
	var maxPairs = len(nums1) * len(nums2)
	// K can be larger than the maximum possible pairs
	// so we need to make sure to exit when no more pairs are
	// available
	for i := 0; i < k && i < maxPairs; i++ {
		// Take smallest pair
		cur := heap.Pop(pq).(Pair)
		res = append(res, []int{cur.A, cur.B})
		// If the 2nd element is not the last one in nums2
		// we can consider a new pair with the next element in nums2
		if cur.Idx < len(nums2)-1 {
			// Push a new pair with the next element in nums2
			// and the smallest element in nums1 so far
			heap.Push(pq, Pair{cur.A, nums2[cur.Idx+1], cur.Idx + 1})
		}
	}
	return res
}

func kSmallestPairsNaive(nums1, nums2 []int, k int) [][]int {
	var res = make([][]int, len(nums1)*len(nums2))
	for i := range nums1 {
		for j := range nums2 {
			res[i*len(nums2)+j] = []int{nums1[i], nums2[j]}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return (res[i][0] + res[i][1]) < (res[j][0] + res[j][1])
	})
	if k >= len(res) {
		return res
	}
	return res[:k]
}

// Priority queue implementation
type Heap struct {
	Values   []Pair
	LessFunc func(Pair, Pair) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() Pair         { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(Pair)) }

func MakeHeap(less func(Pair, Pair) bool) *Heap {
	return &Heap{LessFunc: less}
}

func MaxHeap(i, j Pair) bool { return i.A+i.B > j.A+j.B }
func MinHeap(i, j Pair) bool { return i.A+i.B < j.A+j.B }
