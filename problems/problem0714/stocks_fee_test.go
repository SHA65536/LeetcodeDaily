package problem0714

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Nums     []int
	Fee      int
	Expected int
}

var TestCases = []TestCase{
	{[]int{1, 3, 2, 8, 4, 9}, 2, 8},
	{[]int{1, 3, 7, 5, 10, 3}, 3, 6},
}

func TestProfitStocksWithFee(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maxProfit(tc.Nums, tc.Fee)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkProfitStocksWithFee(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maxProfit(tc.Nums, tc.Fee)
		}
	}
}
