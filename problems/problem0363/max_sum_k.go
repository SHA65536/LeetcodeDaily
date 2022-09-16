package problem0363

/*
Given an m x n matrix matrix and an integer k,
return the max sum of a rectangle in the matrix such that its sum is no larger than k.
It is guaranteed that there will be a rectangle with a sum no larger than k.
*/

const minInt int = -1000000

func maxSumSubmatrix(matrix [][]int, k int) int {
	var maxSum int = minInt
	var curSum int
	// Creating prefix sum matrix
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
	// Searching all squares
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			for g := i; g < len(matrix); g++ {
				for l := j; l < len(matrix[0]); l++ {
					curSum = matrix[g][l]
					if i > 0 {
						curSum -= matrix[i-1][l]
					}
					if j > 0 {
						curSum -= matrix[g][j-1]
					}
					if i > 0 && j > 0 {
						curSum += matrix[i-1][j-1]
					}
					if curSum > maxSum && curSum <= k {
						maxSum = curSum
						if maxSum == k {
							return maxSum
						}
					}
				}
			}
		}
	}
	return maxSum
}
