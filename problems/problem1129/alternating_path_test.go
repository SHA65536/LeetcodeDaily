package problem1129

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N        int
	Reds     [][]int
	Blues    [][]int
	Expected []int
}

var Results = []Result{
	{
		3,
		[][]int{{0, 1}, {1, 2}},
		[][]int{},
		[]int{0, 1, -1},
	},
	{
		3,
		[][]int{{0, 1}},
		[][]int{{2, 1}},
		[]int{0, 1, -1},
	},
	{
		3,
		[][]int{{0, 1}},
		[][]int{{1, 2}},
		[]int{0, 1, 2},
	},
}

func TestShortestAlternatingPath(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := shortestAlternatingPaths(res.N, res.Reds, res.Blues)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
