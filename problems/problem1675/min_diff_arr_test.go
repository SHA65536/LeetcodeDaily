package problem1675

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
	{[]int{1, 2, 3, 4}, 1},
	{[]int{4, 1, 5, 20, 3}, 3},
	{[]int{2, 10, 8}, 3},
}

func TestMinDeviationArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minimumDeviation(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
