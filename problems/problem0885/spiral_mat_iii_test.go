package problem0885

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Rows, Cols int
	SRow, SCol int
	Expected   [][]int
}

var Results = []Result{
	{1, 4, 0, 0, [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
	{5, 6, 1, 4, [][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}}},
}

func TestSpiralMatrixIII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := spiralMatrixIII(res.Rows, res.Cols, res.SRow, res.SCol)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
