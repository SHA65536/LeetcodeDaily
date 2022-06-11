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
	{4, 2, [][]int{}},
}

func TestCombine(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := combine(res.InN, res.InK)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
