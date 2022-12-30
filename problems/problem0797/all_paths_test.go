package problem0797

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected [][]int
}

var Results = []Result{
	{[][]int{{1, 2}, {3}, {3}, {}}, [][]int{{0, 1, 3}, {0, 2, 3}}},
	{[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}},
}

func TestAllPathsToTarget(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := allPathsSourceTarget(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
