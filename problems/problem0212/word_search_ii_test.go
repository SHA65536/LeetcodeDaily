package problem0212

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]byte
	Word     []string
	Expected []string
}

var Results = []Result{
	{
		[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
		[]string{"oath", "pea", "eat", "rain"},
		[]string{"eat", "oath"},
	},
	{
		[][]byte{{'a', 'b'}, {'c', 'd'}},
		[]string{"abcb"},
		[]string{},
	},
}

func TestWordSearchII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findWords(res.Input, res.Word)
		sort.Strings(got)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
