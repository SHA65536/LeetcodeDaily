package problem0373

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Nums1, Nums2 []int
	k            int
	Expected     [][]int
}

var TestCases = []TestCase{
	{[]int{1, 7, 11}, []int{2, 4, 6}, 3,
		[][]int{{1, 2}, {1, 4}, {1, 6}}},
	{[]int{1, 1, 2}, []int{1, 2, 3}, 2,
		[][]int{{1, 1}, {1, 1}}},
	{[]int{1, 2}, []int{3}, 3,
		[][]int{{1, 3}, {2, 3}}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
		6,
		[][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}, {1, 3}, {3, 1}}},
}

func TestSmallestPairSums(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := kSmallestPairs(tc.Nums1, tc.Nums2, tc.k)
		assert.ElementsMatch(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSmallestPairSums(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			kSmallestPairs(tc.Nums1, tc.Nums2, tc.k)
		}
	}
}

func TestSmallestPairSumsNaive(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := kSmallestPairsNaive(tc.Nums1, tc.Nums2, tc.k)
		assert.ElementsMatch(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSmallestPairSumsNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			kSmallestPairsNaive(tc.Nums1, tc.Nums2, tc.k)
		}
	}
}
