package problem0228

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []int
	Expected []string
}

var TestCases = []TestCase{
	{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
	{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
	{[]int{}, []string{}},
	{[]int{1}, []string{"1"}},
	{[]int{1, 2}, []string{"1->2"}},
	{[]int{1, 3}, []string{"1", "3"}},
}

func TestSummaryRanges(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := summaryRanges(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSummaryRanges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			summaryRanges(tc.Input)
		}
	}
}
