package problem0047

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected [][]int
}

var Results = []Result{
	{[]int{1, 1, 2},
		[][]int{{1, 2, 1}, {1, 1, 2}, {2, 1, 1}}},
	{[]int{1, 2, 3},
		[][]int{{1, 3, 2}, {1, 2, 3}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
}

func TestPermuteUnique(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := permuteUnique(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
