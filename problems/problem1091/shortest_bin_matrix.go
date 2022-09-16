package problem1091

/*
https://leetcode.com/problems/shortest-path-in-binary-matrix/

Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. If there is no clear path, return -1.

A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:

All the visited cells of the path are 0.
All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).
The length of a clear path is the number of visited cells of this path
*/

type Point struct {
	Y int
	X int
}

var around = [8]Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	var steps, length = 1, len(grid) - 1
	var visited = map[Point]bool{{0, 0}: true}
	var current = []Point{{0, 0}}
	if grid[0][0] == 1 || grid[length][length] == 1 {
		return -1
	}
	if length == 0 {
		return 1
	}
	//using breadth first search we will
	//check all the routes, until we reach the end
	for len(current) > 0 {
		next := []Point{}
		for _, pt := range current {
			//checking if we arrived
			if pt.X == length && pt.Y == length {
				return steps
			}
			//checking all the points around
			for _, ptD := range around {
				ptN := Point{pt.Y + ptD.Y, pt.X + ptD.X}
				_, ok := visited[ptN]
				if ok || checkInvalid(ptN, length) || grid[ptN.Y][ptN.X] == 1 {
					continue
				}
				visited[ptN] = true
				next = append(next, ptN)
			}
		}
		steps++
		current = next
	}
	return -1
}

func checkInvalid(pt Point, length int) bool {
	return pt.X > length || pt.Y > length || pt.X < 0 || pt.Y < 0
}
