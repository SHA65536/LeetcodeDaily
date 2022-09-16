package problem1354

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
	{[]int{9, 3, 5}, true},
	{[]int{1, 1, 1, 2}, false},
	{[]int{8, 5}, true},
	{[]int{1, 10000000000}, true},
}

func TestIsPossible(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isPossible(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
