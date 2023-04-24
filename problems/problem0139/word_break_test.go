package problem0139

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	S        string
	Dict     []string
	Expected bool
}

var Results = []Result{
	{"leetcode", []string{"leet", "code"}, true},
	{"applepenapple", []string{"apple", "pen"}, true},
	{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
}

func TestWordFromDict(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := wordBreak(res.S, res.Dict)
		assert.Equal(want, got, fmt.Sprintf("%+v", got))
	}
}
