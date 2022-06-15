package problem1048

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected int
}

var Results = []Result{
	{[]string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
	{[]string{"b", "ba", "bda", "bca", "a", "bdca"}, 4},
	{[]string{"bda", "bdca", "a", "b", "ba", "bca"}, 4},
	{[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}, 5},
	{[]string{"pcxbcf", "xbc", "pcxbc", "xb", "cxbc"}, 5},
	{[]string{"cxbc", "xbc", "pcxbc", "pcxbcf", "xb"}, 5},
	{[]string{"cxbc"}, 1},
	{[]string{"cxbc", "apodsa"}, 1},
}

func TestLongestStrChain(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := longestStrChain(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
