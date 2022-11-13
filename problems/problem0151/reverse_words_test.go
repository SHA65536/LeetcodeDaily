package problem0151

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
	{"the sky is blue", "blue is sky the"},
	{"  hello world  ", "world hello"},
	{"a good   example", "example good a"},
}

func TestReverseWords(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reverseWords(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
