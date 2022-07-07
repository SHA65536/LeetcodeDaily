package problem0097

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	S1, S2, S3 string
	Expected   bool
}

var Results = []Result{
	{"aabcc", "dbbca", "aadbbcbcac", true},
	{"aabcc", "dbbca", "aadbbbaccc", false},
	{"", "", "", true},
}

func TestIsInterleave(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isInterleave(res.S1, res.S2, res.S3)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
