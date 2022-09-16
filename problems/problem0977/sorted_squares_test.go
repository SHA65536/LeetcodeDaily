package problem0977

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected []int
}

var Results = []Result{
	{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
}

func TestSortedSquares(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := sortedSquares(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
