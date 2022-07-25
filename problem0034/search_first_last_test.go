package problem0034

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Target   int
	Expected []int
}

var Results = []Result{
	{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	{[]int{5, 6, 7, 7, 8, 8, 10}, 6, []int{1, 1}},
	{[]int{}, 0, []int{-1, -1}},
	{[]int{2, 2}, 3, []int{-1, -1}},
}

func TestSearchRange(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := searchRange(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
