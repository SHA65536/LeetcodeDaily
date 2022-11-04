package problem0345

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
	{"hello", "holle"},
	{"leetcode", "leotcede"},
	{"babebi", "bibeba"},
}

func TestReverseVowelsString(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reverseVowels(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
