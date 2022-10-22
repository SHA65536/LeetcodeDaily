package problem0076

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Target   string
	Expected string
}

var Results = []Result{
	{"ADOBECODEBANC", "ABC", "BANC"},
	{"a", "a", "a"},
	{"a", "aa", ""},
}

func TestMinimumWindow(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minWindow(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
