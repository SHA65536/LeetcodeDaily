package problem0731

/*
You are implementing a program to use as your calendar. We can add a new event if adding the event will not cause a triple booking.
A triple booking happens when three events have some non-empty intersection (i.e., some moment is common to all the three events.).
The event can be represented as a pair of integers start and end that represents a booking on the half-open interval [start, end),
the range of real numbers x such that start <= x < end.
Implement the MyCalendarTwo class:
MyCalendarTwo() Initializes the calendar object.
boolean book(int start, int end) Returns true if the event can be added to the calendar successfully without causing a triple booking.
Otherwise, return false and do not add the event to the calendar.
*/

type MyCalendarTwo struct {
	Calendar [][2]int
	Overlaps [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		Calendar: make([][2]int, 0),
		Overlaps: make([][2]int, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	// Checking if event is already double booked
	for _, event := range this.Overlaps {
		if start < event[1] && end > event[0] {
			return false
		}
	}
	// Checking if event is already booked once
	for _, event := range this.Calendar {
		if start < event[1] && end > event[0] {
			// If it's already booked, add it to overlaps
			this.Overlaps = append(this.Overlaps, [2]int{
				max(start, event[0]),
				min(end, event[1]),
			})
		}
	}
	this.Calendar = append(this.Calendar, [2]int{start, end})
	return true
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
