package problem1328

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
	{"abccba", "aaccba"},
	{"a", ""},
	{"aaaaa", "aaaab"},
}

func TestBreakPalindrome(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := breakPalindrome(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
