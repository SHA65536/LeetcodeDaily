package problem2389

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Queries  []int
	Expected []int
}

var Results = []Result{
	{[]int{4, 5, 2, 1}, []int{3, 10, 21}, []int{2, 3, 4}},
	{[]int{2, 3, 4, 5}, []int{1}, []int{0}},
}

func TestLongestSubsequenceWithSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := answerQueries(res.Input, res.Queries)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
