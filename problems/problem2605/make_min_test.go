package problem2605

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	InA, InB []int
	Expected int
}

var Results = []Result{
	{[]int{4, 1, 3}, []int{5, 7}, 15},
	{[]int{3, 5, 2, 6}, []int{3, 1, 7}, 3},
	{[]int{6, 4, 3, 7, 2, 1, 8, 5}, []int{6, 8, 5, 3, 1, 7, 4, 2}, 1},
}

func TestMakeMinNumber(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minNumber(res.InA, res.InB)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
