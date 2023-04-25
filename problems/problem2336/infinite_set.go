package problem2336

import "container/heap"

/*
You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].
Implement the SmallestInfiniteSet class:
    SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
    int popSmallest() Removes and returns the smallest integer contained in the infinite set.
    void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.
*/

type SmallestInfiniteSet struct {
	Finite    *Heap
	Inifinite int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		Finite:    MakeHeap(MinHeap),
		Inifinite: 1,
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	var res int
	if s.Finite.Len() > 0 {
		res = heap.Pop(s.Finite).(int)
		for s.Finite.Len() > 0 && s.Finite.Peek() == res {
			heap.Pop(s.Finite)
		}
		return res
	}
	s.Inifinite++
	return s.Inifinite - 1
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num < s.Inifinite {
		heap.Push(s.Finite, num)
	}
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
