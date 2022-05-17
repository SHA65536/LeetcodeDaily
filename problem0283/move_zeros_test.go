package problem0283

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected []int
}

var Results = []Result{
	{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	{[]int{0}, []int{0}},
}

func TestMoveZeros(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		moveZeroes(res.Input)
		assert.Equal(want, res.Input, fmt.Sprintf("%+v", res))
	}
}
