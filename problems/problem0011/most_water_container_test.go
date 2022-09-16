package problem0011

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Height   []int
	Expected int
}

/*
Constraints:
	n == height.length
	2 <= n <= 105
	0 <= height[i] <= 104
*/

var Results = []Result{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{[]int{1, 1}, 1},
	{[]int{50, 1, 50}, 100},
}

func TestMaxArea(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		assert.Equal(want, maxArea(res.Height), fmt.Sprintf("%+v", res))
	}
}
