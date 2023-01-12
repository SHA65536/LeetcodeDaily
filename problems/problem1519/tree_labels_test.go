package problem1519

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N        int
	Edges    [][]int
	Lables   string
	Expected []int
}

var Results = []Result{
	{
		7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		"abaedcd", []int{2, 1, 1, 1, 1, 1, 1},
	},
	{
		4, [][]int{{0, 1}, {1, 2}, {0, 3}},
		"bbbb", []int{4, 2, 1, 1},
	},
	{
		5, [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}},
		"aabab", []int{3, 2, 1, 1, 1},
	},
}

func TestSubTreeLabels(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countSubTrees(res.N, res.Edges, res.Lables)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
