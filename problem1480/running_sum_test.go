package problem1480

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
	{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
	{[]int{1, 1, 1, 1}, []int{1, 2, 3, 4}},
}

func TestMergeTwoLists(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := runningSum(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
