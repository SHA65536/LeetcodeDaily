package problem0074

import "sort"

/*
You are given an m x n integer matrix matrix with the following two properties:
    Each row is sorted in non-decreasing order.
    The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.
You must write a solution in O(log(m * n)) time complexity.
*/

func searchMatrix(matrix [][]int, target int) bool {
	for i := len(matrix) - 1; i >= 0; i-- {
		// Skip if it's not possible to find in this row
		if matrix[i][0] > target {
			continue
		}
		// Binary search to find in row
		if idx := sort.SearchInts(matrix[i], target); idx < len(matrix[i]) && matrix[i][idx] == target {
			return true
		}
	}
	return false
}
