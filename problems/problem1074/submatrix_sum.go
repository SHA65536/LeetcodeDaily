package problem1074

/*
https://leetcode.com/problems/number-of-submatrices-that-sum-to-target

Given a matrix and a target, return the number of non-empty submatrices that sum to target.
A submatrix x1, y1, x2, y2 is the set of all cells matrix[x][y] with x1 <= x <= x2 and y1 <= y <= y2.
Two submatrices (x1, y1, x2, y2) and (x1', y1', x2', y2') are different if they have some coordinate that is different: for example, if x1 != x1'.
*/

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	var sum int
	// This matrix is built to aid us in calculating a sub matrix sum
	matrix = ConstructorPrefix(matrix)
	// Calculating each possible submatric
	for x1 := 0; x1 < len(matrix); x1++ {
		for y1 := 0; y1 < len(matrix[0]); y1++ {
			for x2 := x1; x2 < len(matrix); x2++ {
				for y2 := y1; y2 < len(matrix[0]); y2++ {
					if SumRegionPrefix(matrix, x1, y1, x2, y2) == target {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func ConstructorPrefix(matrix [][]int) [][]int {
	// Each position in the matrix will be the sum of
	// the numbers in (0,0,i,j)
	for i := 0; i < len(matrix); i++ {
		var rollingSum int
		for j := 0; j < len(matrix[0]); j++ {
			rollingSum += matrix[i][j]
			matrix[i][j] = rollingSum
			if i > 0 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}
	return matrix
}

func SumRegionPrefix(matrix [][]int, row1 int, col1 int, row2 int, col2 int) int {
	// The sum of numbers in (0,0,row2,col2)
	var sum int = matrix[row2][col2]
	if row1 > 0 {
		// Subtracting the numbers in (0,0,row1-1,col2)
		sum -= matrix[row1-1][col2]
	}
	if col1 > 0 {
		// Subtracting the numbers in (0,0,row2,col1-1)
		sum -= matrix[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		// The area (0,0,row1-1,col1-1) was subtracted twice
		// so we add it once to negate that.
		sum += matrix[row1-1][col1-1]
	}
	return sum
}
