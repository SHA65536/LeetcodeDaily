package problem0044

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input, Pattern string
	Expected       bool
}

var Results = []Result{
	{"aa", "a", false},
	{"aa", "*", true},
	{"cb", "?a", false},
	{"abcde", "a*e", true},
	{"abcdef", "a*e", false},
	{"a", "?*", true},
	{"bbbababbbbabbbbababbaaabbaababbbaabbbaaaabbbaaaabb", "*b********bb*b*bbbbb*ba", false},
	{"", "", true},
	{"ab", "", false},
	{"", "a", false},
	{"", "*", true},
	{"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb", false},
}

func TestMatchWildCard(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isMatch(res.Input, res.Pattern)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
