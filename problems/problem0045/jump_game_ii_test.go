package problem0045

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
	{[]int{2, 3, 1, 1, 4}, 2},
	{[]int{2, 3, 0, 1, 4}, 2},
}

func TestJumpGameII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := jump(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
