package problem1695

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected int
}

var Results = []Result{
	{[]int{1, 1, 4, 2, 3}, 10},
	{[]int{4, 2, 4, 5, 6}, 17},
	{[]int{5, 2, 1, 2, 5, 2, 1, 2, 5}, 8},
	{[]int{1}, 1},
}

func TestMaximumUniqueSubarray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maximumUniqueSubarray(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
