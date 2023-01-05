package problem0452

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
	{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
	{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
	{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
}

func TestMinArrowsToPopBalloons(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findMinArrowShots(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
