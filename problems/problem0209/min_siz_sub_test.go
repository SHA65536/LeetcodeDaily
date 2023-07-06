package problem0209

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Target   int
	Input    []int
	Expected int
}

var TestCases = []TestCase{
	{7, []int{2, 3, 1, 2, 4, 3}, 2},
	{4, []int{1, 4, 4}, 1},
	{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
}

func TestMinSubarraySumTarget(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minSubArrayLen(tc.Target, tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinSubarraySumTarget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minSubArrayLen(tc.Target, tc.Input)
		}
	}
}
