package problem1232

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    [][]int
	Expected bool
}

var TestCases = []TestCase{
	{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}, true},
	{[][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, false},
	{[][]int{{0, 0}, {0, 1}, {0, -1}}, true},
}

func TestIsStraightLine(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := checkStraightLine(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkIsStraightLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			checkStraightLine(tc.Input)
		}
	}
}
