package problem0665

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected bool
}

var Results = []Result{
	{[]int{}, true},
	{[]int{1, 2, 3, 4, 5, 6, 19, 6}, true},
	{[]int{1, 2, 3, 4, 5, 6, 19, 3}, true},
	{[]int{1, 2, 3, 4, 3, 6, 19, 3}, false},
	{[]int{5, 12, 3}, true},
	{[]int{4, 2, 1}, false},
	{[]int{4, 2, 3}, true},
	{[]int{3, 4, 2, 3}, false},
}

func TestCheckPossibility(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		inCopy := make([]int, len(res.Input))
		copy(inCopy, res.Input)
		want := res.Expected
		got := checkPossibility(inCopy)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
