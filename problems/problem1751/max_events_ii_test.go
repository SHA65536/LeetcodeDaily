package problem1751

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    [][]int
	K        int
	Expected int
}

var TestCases = []TestCase{
	{[][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}, 2, 7},
	{[][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}}, 2, 10},
	{[][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}}, 3, 9},
}

func TestMaxValueFromEvents(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maxValue(tc.Input, tc.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxValueFromEvents(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maxValue(tc.Input, tc.K)
		}
	}
}
