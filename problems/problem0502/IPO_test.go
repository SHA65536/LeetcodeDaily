package problem0502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	K, W             int
	Profits, Capital []int
	Expected         int
}

var Results = []Result{
	{
		2, 0,
		[]int{1, 2, 3}, []int{0, 1, 1},
		4,
	},
	{
		3, 0,
		[]int{1, 2, 3}, []int{0, 1, 2},
		6,
	},
	{
		4, 3,
		[]int{4, 3, 6, 5, 4, 23, 22, 40, 49}, []int{1, 2, 4, 5, 9, 13, 15, 39, 41},
		58,
	},
}

func TestIPO(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findMaximizedCapital(res.K, res.W, res.Profits, res.Capital)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
