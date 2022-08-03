package problem0729

/*
You are implementing a program to use as your calendar. We can add a new event if adding the event will not cause a double booking.
A double booking happens when two events have some non-empty intersection (i.e., some moment is common to both events.).
The event can be represented as a pair of integers:
start and end that represents a booking on the half-open interval [start, end), the range of real numbers x such that start <= x < end.
Implement the MyCalendar class:
MyCalendar() Initializes the calendar object.
boolean book(int start, int end) Returns true if the event can be added to the calendar successfully without causing a double booking.
Otherwise, return false and do not add the event to the calendar.
*/

type MyCalendar struct {
	Starts []int
	Ends   []int
}

func Constructor() MyCalendar {
	return MyCalendar{[]int{}, []int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.Starts) == 0 {
		this.Starts = append(this.Starts, start)
		this.Ends = append(this.Ends, end)
		return true
	}
	startPos := searchInsert(this.Starts, start)
	if startPos == -1 {
		return false
	}
	if startPos == len(this.Starts) {
		if start < this.Ends[startPos-1] {
			return false
		}
		this.Starts = append(this.Starts, start)
		this.Ends = append(this.Ends, end)
		return true
	}
	if startPos == 0 {
		if end <= this.Starts[startPos] {
			this.Starts = insert(this.Starts, startPos, start)
			this.Ends = insert(this.Ends, startPos, end)
			return true
		}
		return false
	}
	if start < this.Ends[startPos-1] || end > this.Starts[startPos] {
		return false
	} else {
		this.Starts = insert(this.Starts, startPos, start)
		this.Ends = insert(this.Ends, startPos, end)
		return true
	}
}

func searchInsert(nums []int, target int) int {
	var res, start, end = -1, 0, len(nums) - 1
	for end >= start {
		mid := ((end - start) / 2) + start
		if nums[mid] == target {
			return -1
		} else if nums[mid] > target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			end = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid + 1
			}
			start = mid + 1
		}
	}
	return res
}

func insert(a []int, index int, value int) []int {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
