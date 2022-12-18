package problem0521

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Other    string
	Expected int
}

var Results = []Result{
	{"aaa", "aaa", -1},
	{"aaa", "ccc", 3},
	{"a", "bbb", 3},
	{"aaa", "b", 3},
}

func TestLongestUncommonSubsequence(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findLUSlength(res.Input, res.Other)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
