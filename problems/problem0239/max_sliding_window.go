package problem0239

import "container/heap"

/*
You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right.
You can only see the k numbers in the window.
Each time the sliding window moves right by one position.
Return the max sliding window.
*/

func maxSlidingWindow(nums []int, k int) []int {
	var pq = MakeHeap(MaxHeap)
	var res = make([]int, len(nums)-k+1)

	for i := 0; i < k; i++ {
		heap.Push(pq, [2]int{nums[i], i})
	}

	res[0] = pq.Peek()[0]
	for i := k; i < len(nums); i++ {
		for pq.Len() != 0 && pq.Peek()[1] <= i-k {
			heap.Pop(pq)
		}
		heap.Push(pq, [2]int{nums[i], i})
		res[i-k+1] = pq.Peek()[0]
	}
	return res
}

type Heap struct {
	Values   [][2]int
	LessFunc func([2]int, [2]int) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() [2]int       { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.([2]int)) }

func MakeHeap(less func([2]int, [2]int) bool) *Heap {
	return &Heap{LessFunc: less}
}

func MaxHeap(i, j [2]int) bool { return i[0] > j[0] }
func MinHeap(i, j [2]int) bool { return i[0] < j[0] }
