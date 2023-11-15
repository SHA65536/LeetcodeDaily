package problem1846

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
	{[]int{2, 2, 1, 2, 1}, 2},
	{[]int{100, 1, 1000}, 3},
	{[]int{1, 2, 3, 4, 5}, 5},
}

func TestMaxElementInArr(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var want = tc.Expected
		var got = maximumElementAfterDecrementingAndRearranging(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxElementInArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maximumElementAfterDecrementingAndRearranging(tc.Input)
		}
	}
}
