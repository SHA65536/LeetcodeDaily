package problem2348

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected int64
}

var Results = []Result{
	{[]int{1, 3, 0, 0, 2, 0, 0, 4}, 6},
	{[]int{0, 0, 0, 2, 0, 0}, 9},
	{[]int{2, 10, 2019}, 0},
}

func TestZeroFilledSubArrays(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := zeroFilledSubarray(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
