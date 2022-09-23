package problem1680

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
	{1, 1},
	{3, 27},
	{12, 505379714},
	{10000, 356435599},
	{100000, 757631812},
}

func TestConcatBinaryConsecutive(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := concatenatedBinary(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
