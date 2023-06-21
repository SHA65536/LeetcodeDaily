package problem2448

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Nums, Cost []int
	Expected   int64
}

var TestCases = []TestCase{
	{[]int{1, 3, 5, 2}, []int{2, 3, 1, 14}, 8},
	{[]int{2, 2, 2, 2, 2}, []int{4, 2, 8, 1, 3}, 0},
	{[]int{337, 289, 92, 436, 96, 325, 250, 244, 43, 131, 60, 26, 97, 478, 479, 440, 390, 408, 78, 34, 83, 160, 12, 120, 422, 338, 405, 330, 101, 470, 180, 149, 236, 59, 244, 68, 382, 261, 37, 343, 459, 230, 227, 80, 35, 245, 282, 341, 95, 377, 272, 435, 177, 348, 489, 46, 196, 423, 12, 146, 422, 348, 178, 491, 486, 12, 421, 370, 110, 59, 117, 405, 243, 52, 389, 100, 8, 150, 310, 471, 117, 487, 371, 134, 489, 173, 88, 45, 378, 22, 413, 112, 110, 382, 106, 298, 103, 55, 131, 450}, []int{313, 122, 273, 89, 498, 82, 281, 259, 268, 110, 300, 399, 242, 203, 370, 203, 193, 357, 350, 374, 194, 298, 382, 464, 496, 176, 403, 365, 338, 86, 180, 388, 297, 267, 423, 167, 273, 369, 302, 161, 449, 390, 316, 298, 125, 276, 152, 241, 144, 340, 213, 91, 39, 81, 386, 154, 19, 420, 477, 134, 35, 90, 379, 326, 458, 170, 210, 325, 242, 27, 11, 171, 318, 18, 234, 282, 109, 15, 402, 234, 132, 200, 377, 437, 158, 349, 8, 155, 258, 65, 470, 498, 38, 242, 473, 168, 107, 176, 39, 176}, 3431800},
}

func TestMinCostMakingArrSame(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minCost(tc.Nums, tc.Cost)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinCostMakingArrSame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minCost(tc.Nums, tc.Cost)
		}
	}
}

func TestMinCostMakingArrSameMap(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minCostMap(tc.Nums, tc.Cost)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinCostMakingArrSameMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minCostMap(tc.Nums, tc.Cost)
		}
	}
}
