package problem0378

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Target   int
	Expected int
}

var Results = []Result{
	{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8, 13},
	{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 9, 15},
	{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 7, 13},
	{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 6, 12},
	{[][]int{{-5}}, 1, -5},
}

func TestKthSmallestSortedMat(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := kthSmallest(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
