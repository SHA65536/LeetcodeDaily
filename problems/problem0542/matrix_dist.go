package problem0542

/*
https://leetcode.com/problems/01-matrix/

Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.
*/

func updateMatrix(mat [][]int) [][]int {
	var res = make([][]int, len(mat))
	var count int
	var changed bool = true
	for idx, row := range mat {
		res[idx] = make([]int, len(row))
		for j := range mat[idx] {
			if mat[idx][j] == 0 {
				res[idx][j] = 0
			} else {
				res[idx][j] = -1
			}
		}
	}
	for changed {
		changed = false
		for i := range mat {
			for j := range mat[0] {
				if res[i][j] == count {
					changed = true
					if i+1 < len(res) && res[i+1][j] == -1 {
						res[i+1][j] = count + 1
					}
					if i-1 >= 0 && res[i-1][j] == -1 {
						res[i-1][j] = count + 1
					}
					if j+1 < len(res[0]) && res[i][j+1] == -1 {
						res[i][j+1] = count + 1
					}
					if j-1 >= 0 && res[i][j-1] == -1 {
						res[i][j-1] = count + 1
					}
				}
			}
		}
		count++
	}
	return res
}

func updateMatrixOpt(mat [][]int) [][]int {
	var res = make([][]int, len(mat))
	var count int
	var next, cur [][2]int
	cur = [][2]int{}
	for idx, row := range mat {
		res[idx] = make([]int, len(row))
		for j := range mat[idx] {
			if mat[idx][j] == 0 {
				cur = append(cur, [2]int{idx, j})
				res[idx][j] = 0
			} else {
				res[idx][j] = -1
			}
		}
	}
	for len(cur) > 0 {
		count++
		next = [][2]int{}
		for _, dot := range cur {
			if dot[0]+1 < len(mat) && res[dot[0]+1][dot[1]] == -1 {
				next = append(next, [2]int{dot[0] + 1, dot[1]})
				res[dot[0]+1][dot[1]] = count
			}
			if dot[0]-1 >= 0 && res[dot[0]-1][dot[1]] == -1 {
				next = append(next, [2]int{dot[0] - 1, dot[1]})
				res[dot[0]-1][dot[1]] = count
			}
			if dot[1]+1 < len(mat[0]) && res[dot[0]][dot[1]+1] == -1 {
				next = append(next, [2]int{dot[0], dot[1] + 1})
				res[dot[0]][dot[1]+1] = count
			}
			if dot[1]-1 >= 0 && res[dot[0]][dot[1]-1] == -1 {
				next = append(next, [2]int{dot[0], dot[1] - 1})
				res[dot[0]][dot[1]-1] = count
			}
		}
		cur = next
	}
	return res
}
