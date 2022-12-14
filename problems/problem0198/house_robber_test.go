package problem0198

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
	{[]int{1, 2, 3, 1}, 4},
	{[]int{2, 7, 9, 3, 1}, 12},
	{[]int{7}, 7},
}

func TestHouseRobber(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := rob(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
