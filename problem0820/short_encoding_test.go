package problem0820

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Words    []string
	Expected int
}

var Results = []Result{
	{[]string{"time", "me", "bell"}, 10},
	{[]string{"t"}, 2},
	{[]string{"ctxdic", "c"}, 7},
}

func TestMinimumLengthEncoding(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minimumLengthEncoding(res.Words)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
