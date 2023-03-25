package problem2316

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	N        int
	Expected int64
}

var Results = []Result{
	{[][]int{{0, 1}, {0, 2}, {1, 2}}, 3, 0},
	{[][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}, 7, 14},
}

func TestUnreachablePairsGraph(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countPairs(res.N, res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
