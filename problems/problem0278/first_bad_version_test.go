package problem0278

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected int
}

var Results = []Result{
	{5, 4},
	{1, 1},
	{234, 11},
}

func TestFirstBadVersion(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		badVersionFunc := func(n int) bool {
			return n >= want
		}
		got := firstBadVersion(res.Input, badVersionFunc)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
