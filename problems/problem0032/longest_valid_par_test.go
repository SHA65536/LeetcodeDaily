package problem0032

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
	{"(()", 2},
	{")()())", 4},
	{"", 0},
}

func TestLongestValidParentheses(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := longestValidParentheses(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
