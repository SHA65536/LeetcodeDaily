package problem0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Str      string
	Expected string
}

var Results = []Result{
	{"babad", "bab"},
	{"cbbd", "bb"},
	{"badab", "badab"},
	{"abcpooloopcbd", "bcpooloopcb"},
	{"abccdbaabdcc", "ccdbaabdcc"},
	{"a", "a"},
	{"abbahmm", "abba"},
	{"hmmabba", "abba"},
}

func TestLongestPalindrome(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := longestPalindrome(res.Str)
		assert.Equal(want, got, res.Str)
	}
}
