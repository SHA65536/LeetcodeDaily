package problem1332

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Str      string
	Expected int
}

var Results = []Result{
	{"", 0},
	{"aba", 1},
	{"aababba", 2},
	{"aaaaaaaaaaab", 2},
}

func TestRemovePalindromeSub(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := removePalindromeSub(res.Str)
		assert.Equal(want, got, res.Str)
	}
}
