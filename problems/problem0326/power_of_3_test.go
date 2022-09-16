package problem0326

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected bool
}

var Results = []Result{
	{1, true},
	{2, false},
	{3, true},
	{4, false},
	{5, false},
	{6, false},
	{7, false},
	{8, false},
	{9, true},
	{10, false},
	{12, false},
	{27, true},
	{28, false},
}

func TestIsPowerOfThree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isPowerOfThree(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
