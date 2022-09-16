package problem0746

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected int
}

var Results = []Result{
	{[]int{10, 15, 20}, 15},
	{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
}

func TestMinCostClimbingStairs(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minCostClimbingStairs(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
