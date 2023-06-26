package problem2462

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Cost         []int
	K, Candidate int
	Expected     int64
}

var TestCases = []TestCase{
	{[]int{48}, 1, 1, 48},
	{[]int{20, 10}, 2, 2, 30},
	{[]int{30, 20, 10}, 3, 3, 60},
	{[]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 9, 11},
	{[]int{1, 2, 4, 1}, 3, 3, 4},
	{[]int{57, 33, 26, 76, 14, 67, 24, 90, 72, 37, 30}, 11, 2, 526},
}

func TestTotalCostToHire(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := totalCost(tc.Cost, tc.K, tc.Candidate)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkTotalCostToHire(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			totalCost(tc.Cost, tc.K, tc.Candidate)
		}
	}
}
