package problem1061

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	A, B, Base string
	Expected   string
}

var Results = []Result{
	{"parker", "morris", "parser", "makkek"},
	{"hello", "world", "hold", "hdld"},
	{"leetcode", "programs", "sourcecode", "aauaaaaada"},
}

func TestLexicographicallyEqual(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := smallestEquivalentString(res.A, res.B, res.Base)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
