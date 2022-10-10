package problem0342

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
	{-1, false},
	{2, false},
	{3, false},
	{4, true},
	{5, false},
	{6, false},
	{7, false},
	{8, false},
	{9, false},
	{10, false},
	{11, false},
	{12, false},
	{13, false},
	{14, false},
	{15, false},
	{16, true},
	{17, false},
}

func TestPowerOfFour(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isPowerOfFour(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
