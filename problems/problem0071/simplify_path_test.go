package problem0071

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"/home/", "/home"},
	{"/../", "/"},
	{"/home//foo/", "/home/foo"},
	{"/foo/../bar", "/bar"},
}

func TestSimplifyPath(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := simplifyPath(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
