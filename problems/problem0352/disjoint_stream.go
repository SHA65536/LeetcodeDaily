package problem0352

import "sort"

/*
Given a data stream input of non-negative integers a1, a2, ..., an, summarize the numbers seen so far as a list of disjoint intervals.
Implement the SummaryRanges class:
	SummaryRanges() Initializes the object with an empty stream.
	void addNum(int value) Adds the integer value to the stream.
	int[][] getIntervals() Returns a summary of the integers in the stream currently as a list of disjoint intervals [starti, endi].
	The answer should be sorted by start i.
*/

type SummaryRanges struct {
	Sets [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (s *SummaryRanges) AddNum(val int) {
	if len(s.Sets) == 0 {
		s.Sets = append(s.Sets, []int{val, val})
		return
	}
	var idx = sort.Search(len(s.Sets), func(i int) bool {
		return s.Sets[i][0] >= val
	})
	if idx != 0 && s.Sets[idx-1][1]+1 >= val {
		idx--
	}
	var start, end = val, val
	for idx != len(s.Sets) && val+1 >= s.Sets[idx][0] && val-1 <= s.Sets[idx][1] {
		start = min(start, s.Sets[idx][0])
		end = max(end, s.Sets[idx][1])
		s.Sets = remove(s.Sets, idx)
	}
	s.Sets = insert(s.Sets, idx, []int{start, end})
}

func (s *SummaryRanges) GetIntervals() [][]int {
	var res = make([][]int, len(s.Sets))
	copy(res, s.Sets)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func remove(slice [][]int, s int) [][]int {
	return append(slice[:s], slice[s+1:]...)
}

func insert(a [][]int, index int, value []int) [][]int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
