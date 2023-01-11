package problem1443

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N        int
	Edges    [][]int
	Apples   []bool
	Expected int
}

var Results = []Result{
	{
		7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		[]bool{false, false, true, false, true, true, false}, 8,
	},
	{
		7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		[]bool{false, false, true, false, false, true, false}, 6,
	},
	{
		7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		[]bool{false, false, false, false, false, false, false}, 0,
	},
	{
		4, [][]int{{0, 1}, {1, 2}, {0, 3}},
		[]bool{true, true, false, true}, 4,
	},
}

func TestCollectApples(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minTime(res.N, res.Edges, res.Apples)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
