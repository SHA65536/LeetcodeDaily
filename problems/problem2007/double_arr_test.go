package problem2007

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
	{[]int{1, 3, 4, 2, 6, 8}, []int{1, 3, 4}},
	{[]int{6, 3, 0, 1}, []int{}},
	{[]int{1}, []int{}},
}

func TestDoubleArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findOriginalArray(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
