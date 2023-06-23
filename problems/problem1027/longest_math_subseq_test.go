package problem1027

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
	{[]int{3, 6, 9, 12}, 4},
	{[]int{9, 4, 7, 2, 10}, 3},
	{[]int{20, 1, 15, 3, 10, 5, 8}, 4},
}

func TestLongestArithmeticSubseq(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := longestArithSeqLength(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkLongestArithmeticSubseq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			longestArithSeqLength(tc.Input)
		}
	}
}
