package problem2562

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
	{[]int{7, 52, 2, 4}, 596},
	{[]int{5, 14, 13, 8, 12}, 673},
	{[]int{5, 14, 13, 8, 12, 13}, 2063},
	{[]int{5, 14, 13, 8, 12, 13, 1000}, 53733},
}

func TestConcatinationSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findTheArrayConcVal(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
