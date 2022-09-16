package problem0704

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Target   int
	Expected int
}

var Results = []Result{
	{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
	{[]int{-1, 0, 3, 5, 9, 12, 16}, 9, 4},
	{[]int{-1, 0, 3, 5, 9, 12, 16}, 16, 6},
	{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	{[]int{-1, 0, 3, 5, 9, 12, 16}, -1, 0},
	{[]int{-1, 0}, -1, 0},
	{[]int{-1, 0}, 0, 1},
	{[]int{0}, 0, 0},
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := search(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
