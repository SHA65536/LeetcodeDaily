package problem1845

import "container/heap"

/*
Design a system that manages the reservation state of n seats that are numbered from 1 to n.
Implement the SeatManager class:
SeatManager(int n) Initializes a SeatManager object that will manage n seats numbered from 1 to n.
All seats are initially available.
int reserve() Fetches the smallest-numbered unreserved seat, reserves it, and returns its number.
void unreserve(int seatNumber) Unreserves the seat with the given seatNumber.
*/

type SeatManager struct {
	Seats *Heap
	Last  int
}

func Constructor(n int) SeatManager {
	return SeatManager{
		Seats: MakeHeap(MinHeap),
		Last:  1,
	}
}

func (s *SeatManager) Reserve() int {
	if s.Seats.Len() == 0 {
		heap.Push(s.Seats, s.Last)
		s.Last++
	}
	return heap.Pop(s.Seats).(int)
}

func (s *SeatManager) Unreserve(seatNumber int) {
	heap.Push(s.Seats, seatNumber)
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
