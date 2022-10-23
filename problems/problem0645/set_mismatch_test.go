package problem0645

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
	{[]int{1, 2, 2, 4}, []int{2, 3}},
	{[]int{1, 1}, []int{1, 2}},
	{[]int{4, 2, 2, 1}, []int{2, 3}},
}

func TestSetMismatch(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findErrorNums(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
