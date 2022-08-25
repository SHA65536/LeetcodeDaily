package problem0383

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input, Magazine string
	Expected        bool
}

var Results = []Result{
	{"a", "b", false},
	{"aa", "ab", false},
	{"aa", "aab", true},
	{"aaa", "aa", false},
	{"aaab", "baaa", true},
	{"a", "abbbbbbbbbbbbbbbbbbbbbb", true},
}

func TestRansomNote(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := canConstruct(res.Input, res.Magazine)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
