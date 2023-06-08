package problem1351

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    [][]int
	Expected int
}

var TestCases = []TestCase{
	{[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}, 8},
	{[][]int{{3, 2}, {1, 0}}, 0},
}

func TestCountNegativesInGrid(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := countNegatives(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkCountNegativesInGrid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			countNegatives(tc.Input)
		}
	}
}
