package problem0334

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
	{[]int{1, 2, 3, 4, 5}, true},
	{[]int{5, 4, 3, 2, 1}, false},
	{[]int{2, 1, 5, 0, 4, 6}, true},
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, false},
}

func TestIncreasingTriplet(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := increasingTriplet(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
