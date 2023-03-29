package problem1402

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
	{[]int{-1, -8, 0, 5, -9}, 14},
	{[]int{4, 3, 2}, 20},
	{[]int{-1, -4, -5}, 0},
}

func TestReduceDishes(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxSatisfaction(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
