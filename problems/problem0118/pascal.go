package problem0118

/*
Given an integer numRows, return the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it.
*/

func generate(numRows int) [][]int {
	var results = [][]int{{1}}
	for i := 1; i < numRows; i++ {
		var current = make([]int, i+1)
		for j := 1; j < i; j++ {
			current[j] = results[i-1][j-1] + results[i-1][j]
		}
		current[0] = 1
		current[i] = 1
		results = append(results, current)
	}
	return results
}
