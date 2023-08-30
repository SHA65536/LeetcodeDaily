package problem2366

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	Expected int64
}

var TestCases = []TestCase{
	{[]int{3, 9, 3}, 2},
	{[]int{1, 2, 3, 4, 5}, 0},
}

func TestMinSplitForSorted(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minimumReplacement(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinSplitForSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minimumReplacement(tc.Input)
		}
	}
}
