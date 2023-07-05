package problem1493

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
	{[]int{1, 1, 0, 1}, 3},
	{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
	{[]int{1, 1, 1}, 2},
}

func TestLongest1sSubArray(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := longestSubarray(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkLongest1sSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			longestSubarray(tc.Input)
		}
	}
}
