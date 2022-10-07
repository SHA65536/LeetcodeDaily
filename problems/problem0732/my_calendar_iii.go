package problem0732

import "sort"

/*
A k-booking happens when k events have some non-empty intersection (i.e., there is some time that is common to all k events.)
You are given some events [start, end), after each given event, return an integer k representing the maximum k-booking between all the previous events.
Implement the MyCalendarThree class:
MyCalendarThree() Initializes the object.
int book(int start, int end) Returns an integer k representing the largest integer such that there exists a k-booking in the calendar.
*/

type MyCalendarThree struct {
	// Using a map and a slice to mimic an ordered map
	Events map[int]int
	Order  []int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		Events: make(map[int]int),
		Order:  make([]int, 0),
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	var ongoing, res int
	// Adding event start to events
	if _, ok := this.Events[start]; !ok {
		this.Order = InsertSorted(this.Order, start)
	}
	this.Events[start]++ //When event starts, add 1

	// Adding event end to events
	if _, ok := this.Events[end]; !ok {
		this.Order = InsertSorted(this.Order, end)
	}
	this.Events[end]-- //When event ends, subtract 1

	for _, key := range this.Order {
		ongoing += this.Events[key]
		if ongoing > res {
			res = ongoing
		}
	}
	return res
}

// Inserts a number into a sorted list
func InsertSorted(list []int, val int) []int {
	idx := sort.SearchInts(list, val)
	list = append(list, 0)
	copy(list[idx+1:], list[idx:])
	list[idx] = val
	return list
}
