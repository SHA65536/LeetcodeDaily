package problem0035

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Target   int
	Expected int
}

var Results = []Result{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6, 7}, 5, 2},
	{[]int{1, 3, 5, 6, 7}, 2, 1},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1}, 5, 1},
	{[]int{1}, 0, 0},
}

func TestSearchInsert(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := searchInsert(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

/*
Constraints:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
*/
