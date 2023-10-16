package problem0119

/*
Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
*/

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}

	var cur = make([]int, 1, rowIndex+1)
	var prev = make([]int, 2, rowIndex+1)
	cur[0] = 1
	prev[0] = 1
	prev[1] = 1

	for i := 2; i <= rowIndex+1; i++ {
		cur = cur[:i]
		cur[0] = 1
		for j := 1; j < len(cur)-1; j++ {
			cur[j] = prev[j-1] + prev[j]
		}
		cur[len(cur)-1] = 1
		prev, cur = cur, prev
	}
	return prev
}
