package problem0077

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	InN      int
	InK      int
	Expected [][]int
}

var Results = []Result{
	{4, 2, [][]int{
		{1, 2}, {1, 3}, {2, 3}, {1, 4}, {2, 4}, {3, 4},
	}},
}

func TestCombine(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := combine(res.InN, res.InK)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
