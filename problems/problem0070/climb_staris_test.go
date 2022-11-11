package problem0070

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected int
}

var Results = []Result{
	{2, 2},
	{3, 3},
	{15, 987},
	{45, 1836311903},
}

func TestStairClimb(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := climbStairs(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
