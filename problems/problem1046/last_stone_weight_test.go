package problem1046

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
	{[]int{2, 7, 4, 1, 8, 1}, 1},
	{[]int{1}, 1},
}

func TestLastStoneWeight(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := lastStoneWeight(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
