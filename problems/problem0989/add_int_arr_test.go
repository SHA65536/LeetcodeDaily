package problem0989

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	K        int
	Expected []int
}

var Results = []Result{
	{[]int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}},
	{[]int{2, 7, 4}, 181, []int{4, 5, 5}},
	{[]int{2, 1, 5}, 806, []int{1, 0, 2, 1}},
}

func TestArrayFormAddition(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := addToArrayForm(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
