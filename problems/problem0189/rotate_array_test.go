package problem0189

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Steps    int
	Expected []int
}

var Results = []Result{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	{[]int{-1, -100, 3, 99}, 10, []int{3, 99, -1, -100}},
}

func TestRotateArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		rotate(res.Input, res.Steps)
		assert.Equal(want, res.Input)
	}
}

/*
Constraints:
1 <= nums.length <= 10^5
-2^31 <= nums[i] <= 2^31 - 1
0 <= k <= 10^5
*/
