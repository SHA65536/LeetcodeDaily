package problem0435

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
	{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
	{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
	{[][]int{{1, 2}, {2, 3}}, 0},
}

func TestRemoveOverlapping(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := eraseOverlapIntervals(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkRemoveOverlapping(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			eraseOverlapIntervals(tc.Input)
		}
	}
}
