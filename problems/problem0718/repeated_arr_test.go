package problem0718

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input, Second []int
	Expected      int
}

var Results = []Result{
	{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
	{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
}

func TestRepeatedSubarrayLen(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findLength(res.Input, res.Second)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
