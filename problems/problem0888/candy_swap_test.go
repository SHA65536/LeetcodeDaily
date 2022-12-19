package problem0888

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Alice    []int
	Bob      []int
	Expected []int
}

var Results = []Result{
	{[]int{1, 1}, []int{2, 2}, []int{1, 2}},
	{[]int{1, 2}, []int{2, 3}, []int{1, 2}},
	{[]int{2}, []int{1, 3}, []int{2, 3}},
}

func TestFairCandySwap(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := fairCandySwap(res.Alice, res.Bob)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
