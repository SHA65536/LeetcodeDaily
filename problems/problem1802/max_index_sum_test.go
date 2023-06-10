package problem1802

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N, Idx, Sum int
	Expected    int
}

var TestCases = []TestCase{
	{4, 2, 6, 2},
	{6, 1, 10, 3},
}

func TestMaxIndexInArrayLowerThanSum(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maxValue(tc.N, tc.Idx, tc.Sum)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxIndexInArrayLowerThanSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maxValue(tc.N, tc.Idx, tc.Sum)
		}
	}
}
