package problem0219

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	K        int
	Expected bool
}

var Results = []Result{
	{[]int{1, 2, 3, 1}, 3, true},
	{[]int{1, 0, 1, 1}, 1, true},
	{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	{[]int{1, 2, 3}, 0, false},
}

func TestContainsDuplicateII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := containsNearbyDuplicate(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
