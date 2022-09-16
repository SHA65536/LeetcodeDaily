package problem0031

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
	{[]int{1, 2, 3}, []int{1, 3, 2}},
	{[]int{1, 3, 2}, []int{2, 1, 3}},
	{[]int{2, 3, 1}, []int{3, 1, 2}},
	{[]int{3, 2, 1}, []int{1, 2, 3}},
	{[]int{1, 1, 5}, []int{1, 5, 1}},
	{[]int{1}, []int{1}},
}

func TestNextPermutaion(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		nextPermutation(res.Input)
		assert.Equal(want, res.Input, fmt.Sprintf("%+v", res))
	}
}
