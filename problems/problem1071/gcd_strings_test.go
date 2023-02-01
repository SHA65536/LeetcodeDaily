package problem1071

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	S1, S2   string
	Expected string
}

var Results = []Result{
	{"ABCABC", "ABC", "ABC"},
	{"ABABAB", "ABAB", "AB"},
	{"LEET", "CODE", ""},
}

func TestGcdOfStrings(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		got := gcdOfStrings(res.S1, res.S2)
		want := res.Expected
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
