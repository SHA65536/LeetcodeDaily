package problem2616

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	P        int
	Expected int
}

var TestCases = []TestCase{
	{[]int{10, 1, 2, 7, 1, 3}, 2, 1},
	{[]int{4, 2, 1, 2}, 1, 0},
}

func TestMinPairDiff(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minimizeMax(tc.Input, tc.P)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinPairDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minimizeMax(tc.Input, tc.P)
		}
	}
}
