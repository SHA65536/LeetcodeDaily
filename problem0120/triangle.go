package problem0120

/*
https://leetcode.com/problems/triangle

Given a triangle array, return the minimum path sum from top to bottom.
For each step, you may move to an adjacent number of the row below.
More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.
*/

func minimumTotal(triangle [][]int) int {
	var cache = make(map[[2]int]int)
	return triangleSum(triangle, cache, 0, 0)
}

func triangleSum(triangle [][]int, cache map[[2]int]int, row int, idx int) int {
	var left, right int
	if len(triangle) == row+1 {
		return triangle[row][idx]
	}
	if val, ok := cache[[2]int{row + 1, idx}]; ok {
		left = val
	} else {
		left = triangleSum(triangle, cache, row+1, idx)
	}
	if val, ok := cache[[2]int{row + 1, idx + 1}]; ok {
		right = val
	} else {
		right = triangleSum(triangle, cache, row+1, idx+1)
	}
	if right > left {
		cache[[2]int{row, idx}] = triangle[row][idx] + left
		return triangle[row][idx] + left
	} else {
		cache[[2]int{row, idx}] = triangle[row][idx] + right
		return triangle[row][idx] + right
	}
}
