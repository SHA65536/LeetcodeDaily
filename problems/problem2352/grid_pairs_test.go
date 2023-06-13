package problem2352

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
	{[][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, 1},
	{[][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, 3},
}

func TestColAndRowPairsLoops(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := equalPairsLoops(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkColAndRowPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			equalPairsLoops(tc.Input)
		}
	}
}

func TestColAndRowPairsMap(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := equalPairsMap(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkColAndRowPairsMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			equalPairsMap(tc.Input)
		}
	}
}
