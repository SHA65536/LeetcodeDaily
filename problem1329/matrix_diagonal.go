package problem1329

import "sort"

/*
A matrix diagonal is a diagonal line of cells
starting from some cell in either the topmost row or leftmost column and going in the bottom-right direction until reaching the matrix's end.
For example, the matrix diagonal starting from mat[2][0], where mat is a 6 x 3 matrix, includes cells mat[2][0], mat[3][1], and mat[4][2].
Given an m x n matrix mat of integers, sort each matrix diagonal in ascending order and return the resulting matrix.
*/

func diagonalSort(mat [][]int) [][]int {
	var tempSlice []int
	for r := 0; r < len(mat); r++ {
		tempSlice = make([]int, 0, len(mat))
		for j, i := 0, r; j < len(mat[0]) && i < len(mat); j, i = j+1, i+1 {
			tempSlice = append(tempSlice, mat[i][j])
		}
		sort.Ints(tempSlice)
		for j, i := 0, r; j < len(mat[0]) && i < len(mat); j, i = j+1, i+1 {
			mat[i][j] = tempSlice[j]
		}
	}
	for c := 1; c < len(mat[0]); c++ {
		tempSlice = make([]int, 0, len(mat))
		for i, j := 0, c; j < len(mat[0]) && i < len(mat); j, i = j+1, i+1 {
			tempSlice = append(tempSlice, mat[i][j])
		}
		sort.Ints(tempSlice)
		for i, j := 0, c; j < len(mat[0]) && i < len(mat); j, i = j+1, i+1 {
			mat[i][j] = tempSlice[i]
		}
	}
	return mat
}
