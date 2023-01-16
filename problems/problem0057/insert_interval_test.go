package problem0057

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Intervals [][]int
	Input     []int
	Expected  [][]int
}

var Results = []Result{
	{
		[][]int{{1, 3}, {6, 9}},
		[]int{2, 5},
		[][]int{{1, 5}, {6, 9}},
	},
	{
		[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
		[]int{4, 8},
		[][]int{{1, 2}, {3, 10}, {12, 16}},
	},
}

func TestInsertIntervals(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := insert(res.Intervals, res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
