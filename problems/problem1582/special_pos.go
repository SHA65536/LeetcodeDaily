package problem1582

/*
Given an m x n binary matrix mat, return the number of special positions in mat.
A position (i, j) is called special if mat[i][j] == 1 and all other elements in row i and column j are 0 (rows and columns are 0-indexed).
*/

func numSpecial(mat [][]int) int {
	var res int
	var rows = make([]int, len(mat))
	var cols = make([]int, len(mat[0]))

	for r := range mat {
		for c := range mat[0] {
			rows[r] += mat[r][c]
			cols[c] += mat[r][c]
		}
	}

	for r := range mat {
		for c := range mat[0] {
			if mat[r][c] == 1 && rows[r] == 1 && cols[c] == 1 {
				res++
			}
		}
	}

	return res
}
