package problem1639

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Target   string
	Expected int
}

var Results = []Result{
	{[]string{"acca", "bbbb", "caca"}, "aba", 6},
	{[]string{"abba", "baab"}, "bab", 4},
}

func TestNumberOfWaysFromListToTarget(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numWays(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
