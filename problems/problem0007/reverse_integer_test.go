package problem0007

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected int
}

var Results = []Result{
	{123, 321},
	{-123, -321},
	{120, 21},
	{-120, -21},
	{2147483647, 0},
	{-2147483648, 0},
}

func TestReverse(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reverse(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
