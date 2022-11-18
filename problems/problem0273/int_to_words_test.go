package problem0273

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected string
}

var Results = []Result{
	{123, "One Hundred Twenty Three"},
	{12345, "Twelve Thousand Three Hundred Forty Five"},
	{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
	{2147483647, "Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven"},
	{0, "Zero"},
}

func TestIntegerToWords(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numberToWords(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
