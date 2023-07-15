package problem1218

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	Diff     int
	Expected int
}

var TestCases = []TestCase{
	{[]int{1, 2, 3, 4}, 1, 4},
	{[]int{1, 3, 5, 7}, 1, 1},
	{[]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2, 4},
}

func TestLongestArithmeticSubsequenceWithDiff(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := longestSubsequence(tc.Input, tc.Diff)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkLongestArithmeticSubsequenceWithDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			longestSubsequence(tc.Input, tc.Diff)
		}
	}
}

func TestLongestArithmeticSubsequenceWithDiffNaive(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := longestSubsequenceNaive(tc.Input, tc.Diff)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkLongestArithmeticSubsequenceWithDiffNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			longestSubsequenceNaive(tc.Input, tc.Diff)
		}
	}
}
