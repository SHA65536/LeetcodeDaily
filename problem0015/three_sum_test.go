package problem0015

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected [][]int
}

var Results = []Result{
	{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, 2, -1}}},
	{[]int{}, [][]int{}},
	{[]int{0}, [][]int{}},
	{[]int{-1, 0, 1, 0}, [][]int{{-1, 0, 1}}},
}

func TestRomanToInt(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		assert.Equal(want, threeSum(res.Input), fmt.Sprintf("%+v", res))
	}
}
