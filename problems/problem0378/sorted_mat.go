package problem0378

/*
Given an n x n matrix where each of the rows and columns is sorted in ascending order, return the kth smallest element in the matrix.
Note that it is the kth smallest element in the sorted order, not the kth distinct element.
You must find a solution with a memory complexity better than O(n2).
*/

func kthSmallest(matrix [][]int, k int) int {
	var length = len(matrix)
	var left, right, res int = matrix[0][0], matrix[length-1][length-1], -1
	for left <= right {
		var mid = (left + right) / 2
		if countLess(matrix, length, mid) >= k {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func countLess(matrix [][]int, length, x int) int {
	var count, c = 0, length - 1
	for r := 0; r < length; r++ {
		for c >= 0 && matrix[r][c] > x {
			c--
		}
		count += c + 1
	}
	return count
}
