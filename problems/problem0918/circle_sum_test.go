package problem0918

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
	{[]int{1, -2, 3, -2}, 3},
	{[]int{5, -3, 5}, 10},
	{[]int{-3, -2, -3}, -2},
}

func TestCircleSumMax(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxSubarraySumCircular(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
