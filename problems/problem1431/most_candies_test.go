package problem1431

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Extra    int
	Expected []bool
}

var Results = []Result{
	{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
	{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	{[]int{12, 1, 12}, 10, []bool{true, false, true}},
}

func TestKidsWithMostCandies(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := kidsWithCandies(res.Input, res.Extra)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
