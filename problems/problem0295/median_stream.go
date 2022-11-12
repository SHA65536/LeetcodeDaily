package problem0295

import (
	"container/heap"
	"sort"
)

/*
The median is the middle value in an ordered integer list.
If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.
	For example, for arr = [2,3,4], the median is 3.
	For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
Implement the MedianFinder class:
    MedianFinder() initializes the MedianFinder object.
    void addNum(int num) adds the integer num from the data stream to the data structure.
    double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.
*/

type MedianFinder struct {
	Vals []int
}

func Constructor() MedianFinder {
	return MedianFinder{[]int{}}
}

func (this *MedianFinder) AddNum(num int) {
	idx := sort.Search(len(this.Vals), func(i int) bool {
		return this.Vals[i] >= num
	})
	this.Vals = insert(this.Vals, idx, num)
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.Vals)%2 == 0 {
		return float64(this.Vals[len(this.Vals)/2-1]+this.Vals[len(this.Vals)/2]) / 2
	}
	return float64(this.Vals[len(this.Vals)/2])
}

func insert(a []int, index int, value int) []int {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

type MedianFinderHeap struct {
	Left  *Heap
	Right *Heap
}

func ConstructorHeap() MedianFinderHeap {
	return MedianFinderHeap{
		Left:  MakeHeap(func(i1, i2 int) bool { return i1 > i2 }),
		Right: MakeHeap(func(i1, i2 int) bool { return i1 < i2 }),
	}
}

func (this *MedianFinderHeap) AddNum(num int) {
	if this.Left.Len() == this.Right.Len() {
		heap.Push(this.Left, num)
	} else {
		heap.Push(this.Right, num)
	}

	if this.Right.Len() > 0 && this.Left.Peek() >= this.Right.Peek() {
		this.Right.Values[0], this.Left.Values[0] = this.Left.Values[0], this.Right.Values[0]
		heap.Fix(this.Left, 0)
		heap.Fix(this.Right, 0)
	}
}

func (this *MedianFinderHeap) FindMedian() float64 {
	if this.Left.Len() == this.Right.Len() {
		return float64(this.Left.Peek()+this.Right.Peek()) / 2
	}
	return float64(this.Left.Peek())
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
