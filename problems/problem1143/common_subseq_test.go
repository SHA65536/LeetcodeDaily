package problem1143

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Word1    string
	Word2    string
	Expected int
}

var Results = []Result{
	{"abcde", "ace", 3},
	{"abc", "abc", 3},
	{"abc", "def", 0},
	{"abcde", "abeg", 3},
}

func TestLongestCommonSubsequence(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := longestCommonSubsequence(res.Word1, res.Word2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			longestCommonSubsequence(res.Word1, res.Word2)
		}
	}
}
