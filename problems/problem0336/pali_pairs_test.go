package problem0336

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected [][]int
}

var Results = []Result{
	{[]string{"abcd", "dcba", "lls", "s", "sssll"}, [][]int{{0, 1}, {1, 0}, {2, 4}, {3, 2}}},
	{[]string{"bat", "tab", "cat"}, [][]int{{0, 1}, {1, 0}}},
	{[]string{"a", ""}, [][]int{{0, 1}, {1, 0}}},
	{[]string{"a", "abc", "aba", ""}, [][]int{{0, 3}, {2, 3}, {3, 0}, {3, 2}}},
}

func TestPalindromePairs(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := palindromePairs(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
