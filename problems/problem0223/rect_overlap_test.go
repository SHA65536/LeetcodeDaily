package problem0223

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [8]int
	Expected int
}

var Results = []Result{
	{[8]int{-3, 0, 3, 4, 0, -1, 9, 2}, 45},
	{[8]int{-2, -2, 2, 2, -2, -2, 2, 2}, 16},
}

func TestRectangleOverlap(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := computeArea(res.Input[0], res.Input[1], res.Input[2], res.Input[3], res.Input[4], res.Input[5], res.Input[6], res.Input[7])
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
