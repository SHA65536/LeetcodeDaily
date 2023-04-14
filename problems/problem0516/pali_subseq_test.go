package problem0516

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected int
}

var Results = []Result{
	{"bbbab", 4},
	{"cbbd", 2},
}

func TestPalindromeSubsequence(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := longestPalindromeSubseq(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
