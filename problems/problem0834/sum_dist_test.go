package problem0834

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N        int
	Input    [][]int
	Expected []int
}

var Results = []Result{
	{6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}, []int{8, 12, 6, 10, 10, 10}},
	{1, [][]int{}, []int{0}},
	{2, [][]int{{1, 0}}, []int{1, 1}},
}

func TestSumOfDistancesGraph(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := sumOfDistancesInTree(res.N, res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
