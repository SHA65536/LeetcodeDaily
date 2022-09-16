package problem0890

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Pattern  string
	Expected []string
}

var Results = []Result{
	{[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb", []string{"mee", "aqq"}},
	{[]string{"a", "b", "c"}, "a", []string{"a", "b", "c"}},
}

func TestFindAndReplacePattern(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findAndReplacePattern(res.Input, res.Pattern)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
