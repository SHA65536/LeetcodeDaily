package problem0792

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Words    []string
	Expected int
}

var Results = []Result{
	{"abcde", []string{"a", "bb", "acd", "ace"}, 3},
	{"dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}, 2},
}

func TestNumMatchingSubseq(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numMatchingSubseq(res.Input, res.Words)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
