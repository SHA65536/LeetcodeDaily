package problem0630

import (
	"container/heap"
	"sort"
)

/*
https://leetcode.com/problems/course-schedule-iii/

There are n different online courses numbered from 1 to n.
You are given an array courses where courses[i] = [duration, lastDay] indicate that the ith course should be taken continuously for
duration days and must be finished before or on lastDay.
You will start on the 1st day and you cannot take two or more courses simultaneously.
Return the maximum number of courses that you can take.
*/

func scheduleCourse(courses [][]int) int {
	var times = &MinHeap{}
	var time int
	heap.Init(times)
	// Sorting courses by due date, want to work on the earliest due date first
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	for _, course := range courses {
		// Try to work on the next course
		time += course[0]
		// Remember to use heap.Push(yourheap, x) and not yourheap.Push(x) !!!!
		heap.Push(times, course[0])
		// If we're overdue, dump the most costly course
		if time > course[1] {
			time -= heap.Pop(times).(int)
		}
		if time > course[1] {
			break
		}
	}
	return times.Len()
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] > h[j] }
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
