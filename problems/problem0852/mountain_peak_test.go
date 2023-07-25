package problem0852

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
	{[]int{0, 1, 0}, 1},
	{[]int{0, 2, 1, 0}, 1},
	{[]int{0, 10, 5, 2}, 1},
	{[]int{0, 1, 2, 3, 4, 5, 4}, 5},
	{[]int{3, 4, 5, 4, 3, 2, 1}, 2},
	{[]int{0, 1, 2, 3, 4, 5, 6, 5}, 6},
	{[]int{5, 6, 5, 4, 3, 2, 1, 0}, 1},
	{[]int{0, 1, 2, 3, 2, 1, 0}, 3},
}

func TestMountainPeak(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := peakIndexInMountainArray(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMountainPeak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			peakIndexInMountainArray(tc.Input)
		}
	}
}
