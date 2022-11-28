package problem0392

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Target   string
	Expected bool
}

var Results = []Result{
	{"abc", "ahbgdc", true},
	{"axc", "ahbgdc", false},
}

func TestIsSubsequence(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isSubsequence(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
