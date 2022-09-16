package problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Str      string
	Expected int
}

var Results = []Result{
	{"abcabcbb", 3},  //abc
	{"bbbb", 1},      //b
	{"pwwkew", 3},    //wke
	{"abcdebfgh", 7}, //cdebfgh
	{"abcd", 4},      //abcd
	{"", 0},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := lengthOfLongestSubstring(res.Str)
		assert.Equal(want, got, res.Str)
	}
}
