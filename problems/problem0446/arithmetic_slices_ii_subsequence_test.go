package problem0446

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
	{[]int{2, 4, 6, 8, 10}, 7},
	{[]int{7, 7, 7, 7, 7}, 16},
	{[]int{-2147483648, 2147483647, -2147483648}, 0},
}

func TestArithmeticSubNum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numberOfArithmeticSlices(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
