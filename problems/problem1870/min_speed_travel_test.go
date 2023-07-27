package problem1870

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	Hour     float64
	Expected int
}

var TestCases = []TestCase{
	{[]int{1, 3, 2}, 6, 1},
	{[]int{1, 3, 2}, 2.7, 3},
	{[]int{1, 3, 2}, 1.9, -1},
	{[]int{1, 1}, 1, -1},
}

func TestMinSpeedTravelTime(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minSpeedOnTime(tc.Input, tc.Hour)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinSpeedTravelTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minSpeedOnTime(tc.Input, tc.Hour)
		}
	}
}
