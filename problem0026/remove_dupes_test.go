package problem0026

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input       []int
	ExpectedLen int
	Expected    []int
}

var Results = []Result{
	{[]int{1, 1, 2}, 2, []int{1, 2, 2}},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}},
}

func TestRemoveDuplicates(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		got := removeDuplicates(res.Input)
		assert.Equal(res.Expected, res.Input, fmt.Sprintf("%+v", res))
		assert.Equal(res.ExpectedLen, got, fmt.Sprintf("%+v", res))
	}
}
