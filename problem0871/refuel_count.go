package problem0871

import (
	"container/heap"
)

/*
A car travels from a starting position to a destination which is target miles east of the starting position.
There are gas stations along the way.
The gas stations are represented as an array stations
Where stations[i] = [positioni, fueli] indicates that the ith gas station is positioni miles east of the starting position and has fueli liters of gas.
The car starts with an infinite tank of gas, which initially has startFuel liters of fuel in it.
It uses one liter of gas per one mile that it drives.
When the car reaches a gas station, it may stop and refuel, transferring all the gas from the station into the car.
Return the minimum number of refueling stops the car must make in order to reach its destination. If it cannot reach the destination, return -1.
Note that if the car reaches a gas station with 0 fuel left, the car can still refuel there.
If the car reaches the destination with 0 fuel left, it is still considered to have arrived.
*/

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	var fuel, cnt, prev, dis int
	stations = append(stations, []int{target, 0})
	fuel = startFuel
	var miss = &MaxHeap{}
	heap.Init(miss)
	for _, station := range stations {
		var pos, gas = station[0], station[1]
		dis, prev = pos-prev, pos
		if fuel < dis {
			for !miss.Empty() && fuel < dis {
				fuel += -heap.Pop(miss).(int)
				cnt++
			}
			if fuel < dis {
				return -1
			}
		}
		fuel -= dis
		heap.Push(miss, -gas)
	}
	return cnt
}

type MaxHeap []int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Empty() bool        { return len(*h) == 0 }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
