package problem1013

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
	{[]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}, true},
	{[]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}, false},
	{[]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}, true},
}

func TestDivideArrayIntoThree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := canThreePartsEqualSum(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
