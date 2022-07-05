package problem0128

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
	{[]int{100, 4, 200, 1, 3, 2}, 4},
	{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	{[]int{4, 21, 5, 22, 20, 18, 2, 23, 1, 6, 7, 6, 12, 20, 1, 16, 9, 22, 24, 20, 3, 14, 3, 13, 8}, 9},
}

func TestLongestConsecutive(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := longestConsecutive(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
