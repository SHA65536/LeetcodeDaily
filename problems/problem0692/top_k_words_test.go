package problem0692

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	K        int
	Expected []string
}

var Results = []Result{
	{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
	{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
}

func TestTopKFrequent(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := topKFrequent(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
