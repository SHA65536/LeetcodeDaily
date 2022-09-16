package problem0329

/*
https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

Given an m x n integers matrix, return the length of the longest increasing path in matrix.
From each cell, you can either move in four directions: left, right, up, or down.
You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).
*/

type Point struct {
	X int
	Y int
}

var around = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func longestIncreasingPath(matrix [][]int) int {
	var max, val int
	var lengths = map[Point]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			val = search(lengths, matrix, Point{i, j}, Point{-1, -1})
			if val > max {
				max = val
			}
		}
	}
	return max
}

func search(lengths map[Point]int, matrix [][]int, cur, prev Point) int {
	var length int
	if val, ok := lengths[cur]; ok {
		return val
	}
	for _, dPt := range around {
		nPt := Point{cur.X + dPt.X, cur.Y + dPt.Y}
		if nPt.X >= len(matrix) || nPt.X < 0 || nPt.Y >= len(matrix[0]) || nPt.Y < 0 {
			continue
		}
		if nPt == prev || matrix[nPt.X][nPt.Y] <= matrix[cur.X][cur.Y] {
			continue
		}
		if val, ok := lengths[nPt]; ok {
			if val > length {
				length = val
			}
		} else {
			val = search(lengths, matrix, nPt, cur)
			if val > length {
				length = val
			}
		}
	}
	lengths[cur] = length + 1
	return length + 1
}
