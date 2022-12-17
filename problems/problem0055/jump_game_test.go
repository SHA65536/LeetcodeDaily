package problem0055

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected bool
}

var Results = []Result{
	{[]int{2, 3, 1, 1, 4}, true},
	{[]int{3, 2, 1, 0, 4}, false},
	{[]int{1, 3, 2}, true},
}

func TestJumpGame(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := canJump(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
