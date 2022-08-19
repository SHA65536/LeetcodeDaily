package problem0659

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
	{[]int{1, 2, 3, 3, 4, 5}, true},
	{[]int{1, 2, 3, 3, 4, 4, 5, 5}, true},
	{[]int{1, 2, 3, 4, 4, 5}, false},
}

func TestRemoveHalfArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isPossible(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
