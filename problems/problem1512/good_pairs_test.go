package problem1512

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
	{[]int{1, 2, 3, 1, 1, 3}, 4},
	{[]int{1, 1, 1, 1}, 6},
	{[]int{1, 2, 3}, 0},
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 406},
}

func TestGoodPairs(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := numIdenticalPairs(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkGoodPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			numIdenticalPairs(tc.Input)
		}
	}
}

func TestGoodPairsArr(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := numIdenticalPairsArr(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkGoodPairsArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			numIdenticalPairsArr(tc.Input)
		}
	}
}
