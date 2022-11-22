package problem2000

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Char     byte
	Expected string
}

var Results = []Result{
	{"abcdefd", 'd', "dcbaefd"},
	{"xyxzxe", 'z', "zxyxxe"},
	{"abcd", 'z', "abcd"},
}

func TestReversePrefix(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reversePrefix(res.Input, res.Char)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
