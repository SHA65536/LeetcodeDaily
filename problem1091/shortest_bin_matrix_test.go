package problem1091

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected int
}

var Results = []Result{
	{[][]int{
		{0, 1},
		{1, 0}}, 2},
	{[][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0}}, 4},
	{[][]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 0}}, -1},
	{[][]int{
		{0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 1, 0},
		{0, 0, 1, 0, 1, 0},
		{0, 1, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0}}, 14},
}

func TestShortestPathBinaryMatrix(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := shortestPathBinaryMatrix(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

/*
Constraints:
n == grid.length
n == grid[i].length
1 <= n <= 100
grid[i][j] is 0 or 1
*/
