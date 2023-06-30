package problem1970

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Row, Col int
	Cells    [][]int
	Expected int
}

var TestCases = []TestCase{
	{2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}, 2},
	{2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}, 1},
	{3, 3, [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}, 3},
}

func TestLastDayWeCanCross(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := latestDayToCross(tc.Row, tc.Col, tc.Cells)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkLastDayWeCanCross(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			latestDayToCross(tc.Row, tc.Col, tc.Cells)
		}
	}
}
