package problem1615

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N        int
	Roads    [][]int
	Expected int
}

var TestCases = []TestCase{
	{4, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}, 4},
	{5, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}, 5},
	{8, [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}}, 5},
}

func TestMaxNetworkRank(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maximalNetworkRank(tc.N, tc.Roads)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxNetworkRank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maximalNetworkRank(tc.N, tc.Roads)
		}
	}
}
