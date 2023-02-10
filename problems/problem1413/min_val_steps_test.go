package problem1413

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
	{[]int{-3, 2, -3, 4, 2}, 5},
	{[]int{1, 2}, 1},
	{[]int{1, -2, -3}, 5},
}

func TestMinValueStepByStep(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minStartValue(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
