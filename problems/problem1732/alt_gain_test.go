package problem1732

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
	{[]int{-5, 1, 5, 0, -7}, 1},
	{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
}

func TestMaxAltitude(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := largestAltitude(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkFilterRestaurants(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			largestAltitude(tc.Input)
		}
	}
}
