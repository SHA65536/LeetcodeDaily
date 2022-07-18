package problem1074

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Target   int
	Expected int
}

var Results = []Result{
	{[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0, 4},
	{[][]int{{1, -1}, {-1, 1}}, 0, 5},
	{[][]int{{904}}, 0, 0},
}

func TestNumSubmatrixSumTarget(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := numSubmatrixSumTarget(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
