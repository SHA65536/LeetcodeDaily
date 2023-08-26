package problem0646

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
	{[][]int{{1, 2}, {2, 3}, {3, 4}}, 2},
	{[][]int{{1, 2}, {7, 8}, {4, 5}}, 3},
}

func TestLongestPairChain(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := findLongestChain(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkLongestPairChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			findLongestChain(tc.Input)
		}
	}
}
