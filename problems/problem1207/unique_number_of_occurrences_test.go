package problem1207

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected bool
}

var Results = []Result{
	{[]int{1, 2, 2, 1, 1, 3}, true},
	{[]int{1, 2}, false},
	{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
}

func TestUniqueFrequency(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := uniqueOccurrences(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
