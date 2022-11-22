package problem1417

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"a0b1c2", "0a1b2c"},
	{"leetcode", ""},
	{"1229857369", ""},
	{"aaa111", "1a1a1a"},
	{"aaaa111", "a1a1a1a"},
	{"aaa1111", "1a1a1a1"},
	{"111aaaa", "a1a1a1a"},
	{"a", "a"},
	{"1", "1"},
	{"aa", ""},
	{"11", ""},
	{"a1", "1a"},
	{"1a", "1a"},
}

func TestReformatString(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reformat(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
