package problem0329

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
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1}}, 4},
	{[][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1}}, 4},
	{[][]int{
		{9, 8, 7},
		{2, 1, 6},
		{3, 4, 5}}, 9},
	{[][]int{{1}}, 1},
}

func TestLongestIncreasingPath(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := longestIncreasingPath(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
