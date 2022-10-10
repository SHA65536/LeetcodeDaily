package problem0120

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
	{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
	{[][]int{{-1}, {2, 3}, {1, -1, -3}}, -1},
	{[][]int{{-10}}, -10},
}

func TestMinimumTotalTriangle(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minimumTotal(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
