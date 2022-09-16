package problem1383

import (
	"container/heap"
	"sort"
)

/*
You are given two integers n and k and two integer arrays speed and efficiency both of length n.
There are n engineers numbered from 1 to n. speed[i] and efficiency[i] represent the speed and efficiency of the ith engineer respectively.
Choose at most k different engineers out of the n engineers to form a team with the maximum performance.
The performance of a team is the sum of their engineers' speeds multiplied by the minimum efficiency among their engineers.
Return the maximum performance of this team. Since the answer can be a huge number, return it modulo 10^9 + 7.
*/

const mod = 1000000007

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	var total, best int
	// Joining the lists together
	var joined = make([][2]int, n)
	for i := range joined {
		joined[i] = [2]int{efficiency[i], speed[i]}
	}
	// Sort by efficiency
	sort.Slice(joined, func(i, j int) bool {
		return joined[i][0] > joined[j][0]
	})
	// Making our min priority queue
	var speedPq = &MinHeap{}
	heap.Init(speedPq)

	for i := range joined {
		// Making sure we dont have more than k engineers
		if speedPq.Len() >= k {
			total -= heap.Pop(speedPq).(int)
		}
		// Adding a new engineer
		total += joined[i][1]
		heap.Push(speedPq, joined[i][1])
		// Checkign if current performance is best
		if total*joined[i][0] > best {
			best = total * joined[i][0]
		}
	}
	// Answer could be very large so we return mod 10^9+7
	return best % mod
}

// Regular min heap implementation I copied
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}
