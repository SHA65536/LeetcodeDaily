package problem1658

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	X        int
	Expected int
}

var Results = []Result{
	{[]int{1, 1, 4, 2, 3}, 5, 2},
	{[]int{5, 6, 7, 8, 9}, 4, -1},
	{[]int{3, 2, 20, 1, 1, 3}, 10, 5},
}

func TestMinOperations(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minOperations(res.Input, res.X)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
