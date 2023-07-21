package problem0673

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	Expected int
}

var TestCases = []TestCase{
	{[]int{1, 3, 5, 4, 7}, 2},
	{[]int{2, 2, 2, 2, 2}, 5},
}

func TestNumberOfLongestIncreasingSubsequences(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := findNumberOfLIS(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkNumberOfLongestIncreasingSubsequences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			findNumberOfLIS(tc.Input)
		}
	}
}
