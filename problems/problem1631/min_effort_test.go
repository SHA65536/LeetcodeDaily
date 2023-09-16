package problem1631

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
	{[][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}, 2},
	{[][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}, 1},
	{[][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}, 0},
}

func TestMinEffort(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minimumEffortPath(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinEffort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minimumEffortPath(tc.Input)
		}
	}
}
