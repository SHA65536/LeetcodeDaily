package problem0295

import "sort"

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
