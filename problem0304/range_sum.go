package problem0304

/*
https://leetcode.com/problems/range-sum-query-2d-immutable/

Given a 2D matrix matrix, handle multiple queries of the following type:
Calculate the sum of the elements of matrix inside the rectangle defined by its
upper left corner (row1, col1) and lower right corner (row2, col2).

Implement the NumMatrix class:
NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.
int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the elements of matrix inside the rectangle defined by its
upper left corner (row1, col1) and lower right corner (row2, col2).
*/

type NumMatrix struct {
	Matrix [][]int
	Cache  map[[4]int]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{matrix, make(map[[4]int]int)}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	if val, ok := this.Cache[[4]int{row1, col1, row2, col2}]; ok {
		return val
	}
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += this.Matrix[i][j]
		}
	}
	this.Cache[[4]int{row1, col1, row2, col2}] = sum
	return sum
}
