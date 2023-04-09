package problem1857

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Colors   string
	Edges    [][]int
	Expected int
}

var Results = []Result{
	{"abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}, 3},
	{"a", [][]int{{0, 0}}, -1},
}

func TestColorGraph(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := largestPathValue(res.Colors, res.Edges)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
