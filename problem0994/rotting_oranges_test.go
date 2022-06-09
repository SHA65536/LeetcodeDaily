package problem0994

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected int
}

var Results = []Result{
	{
		[][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}, 4,
	},
	{
		[][]int{
			{2, 1, 1},
			{0, 1, 1},
			{1, 0, 1},
		}, -1,
	},
	{
		[][]int{
			{0, 2},
		}, 0,
	},
	{
		[][]int{
			{2, 2},
			{1, 1},
			{0, 0},
			{2, 0},
		}, 1,
	},
}

func TestOrangesRotting(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := orangesRotting(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
