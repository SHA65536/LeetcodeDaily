package problem0823

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected int
}

var Results = []Result{
	{[]int{2, 4}, 3},
	{[]int{2, 4, 5, 10}, 7},
}

func TestNumFactoredBinaryTrees(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		got := numFactoredBinaryTrees(res.Input)
		want := res.Expected
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
