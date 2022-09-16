package problem0016

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Nums     []int
	Target   int
	Expected int
}

var Results = []Result{
	{[]int{-1, 2, 1, -4}, 1, 2}, // (-1 + 2 + 1 = 2)
	{[]int{0, 0, 0}, 1, 0},      // (0 + 0 + 0 = 0)
	{[]int{0, 1, 2}, 0, 3},      // (0 + 1 + 2 = 3)
	{[]int{0, 1, 2}, 4, 3},      // (0 + 1 + 2 = 3)
}

func TestThreeSumClosest(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := threeSumClosest(res.Nums, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

/*
Constraints:
3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-10^4 <= target <= 10^4
*/
