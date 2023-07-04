package problem0137

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
	{[]int{2, 2, 3, 2}, 3},
	{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
}

func TestSingleNumberII(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := singleNumber(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSingleNumberII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			singleNumber(tc.Input)
		}
	}
}
