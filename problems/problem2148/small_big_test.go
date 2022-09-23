package problem2148

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
	{[]int{11, 7, 2, 15}, 2},
	{[]int{-3, 3, 3, 90}, 2},
	{[]int{3, 3, 3, 3}, 0},
	{[]int{3, 3, 4, 5}, 1},
	{[]int{3, 3, 4, 4}, 0},
}

func TestCountBiggerAndSmaller(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countElements(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
