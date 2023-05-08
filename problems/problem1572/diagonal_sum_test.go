package problem1572

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected int
}

var Results = []Result{
	{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 25},
	{[][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, 8},
}

func TestDiagonalSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := diagonalSum(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkDiagonalSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			diagonalSum(res.Input)
		}
	}
}
