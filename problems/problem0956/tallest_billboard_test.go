package problem0956

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
	{[]int{1, 2, 3, 6}, 6},
	{[]int{1, 2, 3, 4, 5, 6}, 10},
	{[]int{1, 2}, 0},
	{[]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 50, 50, 50, 150, 150, 150, 100, 100, 100, 123}, 1023},
}

func TestTallestBillboard(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := tallestBillboard(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkTallestBillboard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			tallestBillboard(tc.Input)
		}
	}
}

func TestTallestBillboardNaive(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := tallestBillboardNaive(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkTallestNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			tallestBillboardNaive(tc.Input)
		}
	}
}
