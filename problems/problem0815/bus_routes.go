package problem0815

import "container/heap"

/*
You are given an array routes representing bus routes where routes[i] is a bus route that the ith bus repeats forever.
For example, if routes[0] = [1, 5, 7], this means that the 0th bus travels in the sequence 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... forever.
You will start at the bus stop source (You are not on any bus initially), and you want to go to the bus stop target.
You can travel between bus stops by buses only.
Return the least number of buses you must take to travel from source to target. Return -1 if it is not possible.
*/

type Pair struct {
	Stop bool // True if it's a stop, false if a bus
	Dst  int  // Destination
	Cost int  // Cost so far
}

func numBusesToDestination(busToStop [][]int, src, dst int) int {
	// busToStop shows which busses go to which stops
	// stopToBus shows which stops go to which busses
	var stopToBus = map[int][]int{}
	// Building graph
	for i := range busToStop {
		for _, s := range busToStop[i] {
			stopToBus[s] = append(stopToBus[s], i)
		}
	}

	// Keep track of visited locations
	var seenStop = map[int]bool{}
	var seenBus = map[int]bool{}

	// MinHeap storing best possible moves so far
	var queue = MakeHeap(MinHeap)
	queue.Push(Pair{Stop: true, Dst: src, Cost: 0})

	// Loop until nothing left to explore
	for queue.Len() > 0 {
		var cur = heap.Pop(queue).(Pair)
		// Check if we've arrived to the destination
		if cur.Stop && cur.Dst == dst {
			return cur.Cost
		}
		if cur.Stop {
			// Make sure we didn't visit this stop yet
			if seenStop[cur.Dst] {
				continue
			}
			seenStop[cur.Dst] = true
			// Add all the busses that can be reached from this stop to the heap
			for _, bus := range stopToBus[cur.Dst] {
				if !seenBus[bus] {
					// Since we entered a bus, we add 1 to the cost
					heap.Push(queue, Pair{Stop: false, Dst: bus, Cost: cur.Cost + 1})
				}
			}
		} else {
			// Make sure we didn't visit this bus yet
			if seenBus[cur.Dst] {
				continue
			}
			seenBus[cur.Dst] = true
			// Add all the stops that this bus can reach to the heap
			for _, stop := range busToStop[cur.Dst] {
				if !seenStop[stop] {
					// Here we don't add cost because we paid when we got on the bus
					heap.Push(queue, Pair{Stop: true, Dst: stop, Cost: cur.Cost})
				}
			}
		}
	}

	return -1
}

// Generic heap implementation
type Heap struct {
	Values   []Pair
	LessFunc func(int, int) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i].Cost, h.Values[j].Cost) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() Pair         { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(Pair)) }

func MakeHeap(less func(int, int) bool) *Heap {
	return &Heap{LessFunc: less}
}

func MaxHeap(i, j int) bool { return i > j }
func MinHeap(i, j int) bool { return i < j }
