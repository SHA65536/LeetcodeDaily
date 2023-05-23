package problem0703

import "container/heap"

/*
Design a class to find the kth largest element in a stream.
Note that it is the kth largest element in the sorted order, not the kth distinct element.
Implement KthLargest class:
KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.
*/

type KthLargest struct {
	nums *Heap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{nums: MakeHeap(MinHeap), k: k}
	for i := range nums {
		kth.Add(nums[i])
	}
	return kth
}

func (kth *KthLargest) Add(val int) int {
	heap.Push(kth.nums, val)
	if kth.nums.Len() > kth.k {
		heap.Pop(kth.nums)
	}
	return kth.nums.Peek()
}

// Standard heap implementation
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
