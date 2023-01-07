package problem0138

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Gas, Cost []int
	Expected  int
}

var Results = []Result{
	{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
	{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
}

func TestCompleteGasCircuit(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := canCompleteCircuit(res.Gas, res.Cost)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
