package problem0658

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	K, X     int
	Expected []int
}

var Results = []Result{
	{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, 4, 10, []int{2, 3, 4, 5}},
	{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5, []int{3, 3, 4}},
	{[]int{1}, 1, 2, []int{1}},
	{[]int{1}, 1, 1, []int{1}},
	{[]int{1}, 1, 0, []int{1}},
}

func TestFindClosestIntsInArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findClosestElements(res.Input, res.K, res.X)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
