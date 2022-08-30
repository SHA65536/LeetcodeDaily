package problem0048

/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
DO NOT allocate another 2D matrix and do the rotation.
*/

func rotate(matrix [][]int) {
	var length = len(matrix)
	for row := 0; row < length/2; row++ {
		for col := row; col < length-row-1; col++ {
			var temp1, temp2 int
			temp1 = matrix[col][length-1-row]
			matrix[col][length-1-row] = matrix[row][col]
			temp2 = matrix[length-1-row][length-1-col]
			matrix[length-1-row][length-1-col] = temp1
			temp1 = matrix[length-1-col][row]
			matrix[length-1-col][row] = temp2
			matrix[row][col] = temp1
		}
	}
}

/*
/-----------------------\
| 0,0 | 0,1 | 0,2 | 0,3 |
|-----------------------|
| 1,0 | 1,1 | 1,2 | 1,3 |
|-----------------------|
| 2,0 | 2,1 | 2,2 | 2,3 |
|-----------------------|
| 3,0 | 3,1 | 3,2 | 3,3 |
\-----------------------/
*/

// 0,0 -> 0,3 -> 3,3 -> 3,0
// 0,1 -> 1,3 -> 3,2 -> 2,0
// 0,2 -> 2,3 -> 3,1 -> 1,0

// 1,1 -> 1,2 -> 2,2 -> 2,1
