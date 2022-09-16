package problem0027

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input       []int
	Val         int
	ExpectedLen int
	Expected    []int
}

var Results = []Result{
	{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2, 2, 3}},
	{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4, 0, 4, 2}},
}

func TestRemoveElement(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		got := removeElement(res.Input, res.Val)
		assert.Equal(res.Expected, res.Input, fmt.Sprintf("%+v", res))
		assert.Equal(res.ExpectedLen, got, fmt.Sprintf("%+v", res))
	}
}
