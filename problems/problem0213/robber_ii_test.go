package problem0213

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
	{[]int{2, 3, 2}, 3},
	{[]int{1, 2, 3, 1}, 4},
	{[]int{1, 2, 3}, 3},
}

func TestHouseRobberII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := rob(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
