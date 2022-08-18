package problem1338

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
	{[]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, 2},
	{[]int{7, 7, 7, 7, 7, 7}, 1},
	{[]int{1, 2, 3, 4, 5, 6}, 3},
	{[]int{7, 7, 2, 3, 1, 5, 8, 1, 6, 1, 9, 3, 9, 7, 10, 10, 2, 5, 9, 6, 4, 4, 1, 10, 5, 8, 9, 6, 4, 6, 7, 1, 7, 7, 8, 10, 9, 1, 6, 9, 8, 4, 4, 8, 5, 9, 7, 3, 3, 8, 2, 1, 9, 6, 10, 8, 4, 10, 1, 9, 4, 1, 4, 8, 3, 5, 1, 9, 9, 6, 1, 6, 8, 4, 10, 1, 9, 2, 2, 8, 4, 6, 4, 1, 10, 6, 8, 10, 6, 4, 9, 6, 5, 7, 7, 5, 1, 1, 8, 8}, 4},
}

func TestRemoveHalfArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minSetSize(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
