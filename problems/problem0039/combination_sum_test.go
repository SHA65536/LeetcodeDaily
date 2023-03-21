package problem0039

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Target   int
	Expected [][]int
}

var Results = []Result{
	{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
	{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	{[]int{2}, 1, nil},
}

func TestCombinationSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := combinationSum(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
