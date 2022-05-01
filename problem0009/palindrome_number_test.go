package problem0009

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
	{10, false},
	{1010, false},
	{-110, false},
	{-101, false},
	{-1, false},
	{1, true},
	{101, true},
	{0, true},
	{10101, true},
	{951159, true},
	{9515159, true},
}

func TestMyAtoi(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isPalindrome(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
