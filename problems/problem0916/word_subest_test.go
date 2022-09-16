package problem0916

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Words    []string
	Expected []string
}

var Results = []Result{
	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"e", "o"},
		[]string{"facebook", "google", "leetcode"},
	},
	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"l", "e"},
		[]string{"apple", "google", "leetcode"},
	},
}

func TestWordSubsets(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := wordSubsets(res.Input, res.Words)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
