package problem0131

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected [][]string
}

var Results = []Result{
	{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
	{"a", [][]string{{"a"}}},
	{"aaa", [][]string{{"a", "a", "a"}, {"a", "aa"}, {"aa", "a"}, {"aaa"}}},
}

func TestPartitionPalindrome(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := partition(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
