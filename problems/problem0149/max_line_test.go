package problem0149

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
	{[][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
	{[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
	{[][]int{{69, 69}}, 1},
}

func TestMaxPointOnALine(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxPoints(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
