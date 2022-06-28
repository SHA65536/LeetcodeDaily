package problem1647

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected int
}

var Results = []Result{
	{"aab", 0},
	{"aaabbbcc", 2},
	{"ceabaacb", 2},
}

func TestMinDeletions(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minDeletions(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
