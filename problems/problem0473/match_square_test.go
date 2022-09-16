package problem0473

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
	{[]int{1, 1, 2, 2, 2}, true},
	{[]int{3, 3, 3, 3, 4}, false},
}

func TestMakesquare(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := makesquare(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
