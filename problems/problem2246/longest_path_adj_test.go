package problem2246

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Labels   string
	Expected int
}

var Results = []Result{
	{[]int{-1, 0, 0, 1, 1, 2}, "abacbe", 3},
	{[]int{-1, 0, 0, 0}, "aabc", 3},
	{[]int{-1, 0, 1}, "aab", 2},
}

func TestLongestPathWithAdjacentLabels(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := longestPath(res.Input, res.Labels)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
