package problem2090

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
	{[]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3, []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}},
	{[]int{100000}, 0, []int{100000}},
	{[]int{100000}, 8, []int{-1}},
}

func TestRadiusKAverage(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := getAverages(tc.Input, tc.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkRadiusKAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			getAverages(tc.Input, tc.K)
		}
	}
}
