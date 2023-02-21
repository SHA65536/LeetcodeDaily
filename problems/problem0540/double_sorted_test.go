package problem0540

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
	{[]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6}, 4},
	{[]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6, 6}, 5},
	{[]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7}, 7},
	{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
	{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	{[]int{1}, 1},
	{[]int{1, 1, 3}, 3},
	{[]int{1, 3, 3}, 1},
}

func TestAddStrings(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := singleNonDuplicate(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
