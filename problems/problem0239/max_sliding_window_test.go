package problem0239

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	K        int
	Expected []int
}

var TestCases = []TestCase{
	{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
	{[]int{1}, 1, []int{1}},
}

func TestMaxSlidingWindow(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maxSlidingWindow(tc.Input, tc.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxSlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maxSlidingWindow(tc.Input, tc.K)
		}
	}
}
