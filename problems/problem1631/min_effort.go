package problem1631

import (
	"container/heap"
	"math"
)

/*
You are a hiker preparing for an upcoming hike.
You are given heights, a 2D array of size rows x columns, where heights[row][col] represents the height of cell (row, col).
You are situated in the top-left cell, (0, 0), and you hope to travel to the bottom-right cell, (rows-1, columns-1) (i.e., 0-indexed).
You can move up, down, left, or right, and you wish to find a route that requires the minimum effort.
A route's effort is the maximum absolute difference in heights between two consecutive cells of the route.
Return the minimum effort required to travel from the top-left cell to the bottom-right cell.
*/

var moves = [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func minimumEffortPath(heights [][]int) int {
	var rows, cols = len(heights), len(heights[0])
	// queue holding the best path for djikstra's
	var queue = MakeHeap(MinHeap)

	// effort[x][y] holds the lowest effort we found using this tile
	var effort = make([][]int, rows)
	for i := range effort {
		effort[i] = make([]int, cols)
		for j := range effort[i] {
			effort[i][j] = math.MaxInt
		}
	}
	effort[0][0] = 0
	heap.Push(queue, Option{0, 0, 0})

	var cur Option
	for queue.Len() > 0 {
		cur = heap.Pop(queue).(Option)
		// Already found a better way with this tile
		if cur.E > effort[cur.X][cur.Y] {
			continue
		}

		// Got to the target
		if cur.X == rows-1 && cur.Y == cols-1 {
			return cur.E
		}

		// Try all move options
		for _, dif := range moves {
			nX, nY := cur.X+dif[0], cur.Y+dif[1]
			// Out of bounds
			if 0 > nX || nX >= rows || 0 > nY || nY >= cols {
				continue
			}
			newE := max(cur.E, abs(heights[nX][nY]-heights[cur.X][cur.Y]))
			// If we found a better effort using this tile, add to queue
			if effort[nX][nY] > newE {
				effort[nX][nY] = newE
				heap.Push(queue, Option{newE, nX, nY})
			}
		}
	}

	return -1
}

type Option struct {
	E, X, Y int // Effort, X, Y
}

// Min heap interface
type Heap struct {
	Values   []Option
	LessFunc func(int, int) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i].E, h.Values[j].E) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() Option       { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(Option)) }

func MakeHeap(less func(int, int) bool) *Heap {
	return &Heap{LessFunc: less}
}

func MaxHeap(i, j int) bool { return i > j }
func MinHeap(i, j int) bool { return i < j }

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}
