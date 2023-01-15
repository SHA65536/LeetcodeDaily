package problem2421

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Vals     []int
	Edges    [][]int
	Expected int
}

var Results = []Result{
	{[]int{1, 3, 2, 1, 3}, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, 6},
	{[]int{1, 1, 2, 2, 3}, [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}, 7},
	{[]int{1}, [][]int{}, 1},
}

func TestAllGoodPaths(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numberOfGoodPaths(res.Vals, res.Edges)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
