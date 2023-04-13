package problem0946

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Pushes   []int
	Pops     []int
	Expected bool
}

var Results = []Result{
	{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
	{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, true},
	{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, true},
}

func TestValidateStack(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := validateStackSequences(res.Pushes, res.Pops)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
