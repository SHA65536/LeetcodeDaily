package problem2256

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
	{[]int{2, 5, 3, 9, 5, 3}, 3},
	{[]int{0}, 0},
	{[]int{4, 2, 0}, 2},
}

func TestMinAvgDiff(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minimumAverageDifference(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
