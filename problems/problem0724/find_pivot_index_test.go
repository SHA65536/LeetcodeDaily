package problem0724

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
	{[]int{1, 7, 3, 6, 5, 6}, 3},
	{[]int{1, 2, 3}, -1},
	{[]int{2, 1, -1}, 0},
}

func TestFindFirstPivot(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := pivotIndex(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
