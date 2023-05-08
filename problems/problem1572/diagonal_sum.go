package problem1572

/*
Given a square matrix mat, return the sum of the matrix diagonals.
Only include the sum of all the elements on the primary diagonal and all the elements on the secondary
diagonal that are not part of the primary diagonal.
*/

func diagonalSum(mat [][]int) int {
	var sum int
	for i := range mat {
		sum += mat[i][i]
		sum += mat[i][len(mat)-i-1]
	}
	if len(mat)%2 != 0 {
		sum -= mat[len(mat)/2][len(mat)/2]
	}
	return sum
}
