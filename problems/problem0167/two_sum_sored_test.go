package problem0167

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
	{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	{[]int{2, 3, 4}, 6, []int{1, 3}},
	{[]int{-1, 0}, -1, []int{1, 2}},
}

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := twoSum(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
