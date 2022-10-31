package problem0766

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected bool
}

var Results = []Result{
	{[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}, true},
	{[][]int{{1}}, true},
	{[][]int{{1, 2}, {2, 2}}, false},
}

func TestShortestPathRemoveObstacles(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isToeplitzMatrix(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
