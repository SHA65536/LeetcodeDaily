package problem1466

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	N        int
	Expected int
}

var Results = []Result{
	{[][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}, 6, 3},
	{[][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}, 5, 2},
	{[][]int{{1, 0}, {2, 0}}, 3, 0},
}

func TestNumberOfRoadChanges(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minReorder(res.N, res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
