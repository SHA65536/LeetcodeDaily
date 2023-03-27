package problem0064

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
	{[][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}, 7},
	{[][]int{
		{1, 2, 3},
		{4, 5, 6},
	}, 12},
}

func TestMinPathSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minPathSum(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
