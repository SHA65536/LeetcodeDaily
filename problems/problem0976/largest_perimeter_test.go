package problem0976

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
	{[]int{2, 1, 2}, 5},
	{[]int{1, 2, 1}, 0},
}

func TestLargestTrianglePerimeter(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := largestPerimeter(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
