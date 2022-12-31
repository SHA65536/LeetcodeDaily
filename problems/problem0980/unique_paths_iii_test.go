package problem0980

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Grid     [][]int
	Expected int
}

var Results = []Result{
	{[][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}, 2},
	{[][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2},
	}, 4},
	{[][]int{
		{0, 1},
		{2, 0},
	}, 0},
}

func TestUniquePathsIII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := uniquePathsIII(res.Grid)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
