package problem0010

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Str      string
	Pat      string
	Expected bool
}

/*
Constraints:
   1 <= s.length <= 20
   1 <= p.length <= 30
   s contains only lowercase English letters.
   p contains only lowercase English letters, '.', and '*'.
   It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
*/

var Results = []Result{
	{"aa", "a", false},
	{"aa", "aa", true},
	{"aab", "aac", false},
	{"aabca", "aabca", true},
	/*{"aa", "a*", true},
	{"ab", ".b", true},
	{"ab", ".*", true},
	{"aaaa", "aa*", true},
	{"abc", "a", false},
	{"abbd", "ab*d", true},
	{"acd", "ab*d", true},
	{"abbcccd", "acd", false},
	{"a", "aaa*", false},*/
}

func TestIsMatch(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isMatch(res.Str, res.Pat)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
