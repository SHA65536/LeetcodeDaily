package problem0066

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
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	{[]int{9}, []int{1, 0}},
	{[]int{8, 9, 9}, []int{9, 0, 0}},
}

func TestPlusOne(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := plusOne(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
