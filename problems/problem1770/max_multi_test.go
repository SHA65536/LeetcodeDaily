package problem1770

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Nums, Multi []int
	Expected    int
}

var Results = []Result{
	{[]int{1, 2, 3}, []int{3, 2, 1}, 14},
	{[]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}, 102},
	{[]int{-59, -14, 95, 17, -52, -14, -2, 30, -88, -9, 26, 17, 25, -51, -35, 75, 17, -16, 17, -30, 3, 15, -67, 83, 8, 94, 51, 66, -91, 63, 30, -15, 95, 18, 36, 13, 14, 93, -24, -44, 23, -46, 38, 58, 19, 27, 9, 57, -36, -68, 7, 54, -89, -30, 20, 100, -37, -69, 20, -19, 62, 95, -31, -98, -56, 88, 83, -37, 57, -33, 72, -96, 6, -9, -14, 24, 63, -99, 60, -54, 43, 44, 87, -32, 58, -42, -68, 98, -22, -17, -93, -46, -12, -63, 93, 43, 5, 55, -69, -56, -84, 87, 6, 29, 82, 51, 92, -88, 36, -86, 40, -76, 73, 80, 66, 73, -24, -15, 45, -81, -27, -14, -4, 4, 45, -69, 37, 41, -25, 96, 21, -67, -36, -6, 5, -45, -24, 12, 47, -43, -92, 8, 27, -49, 12, -10, 26, 76, 98, -24, 10, 59, 46, -17, -23, -91, 47, 28, 79, 19, -1, -99, 42, -16, 87, 73, -64, 10, -14, -81, 28, -25, 82, 76, -93, -100, -81, -83, 61, 71, 41, 72, -33, -63, 37, 56, -20, 25, -45, 84, 44, 93, 29, 16, 15, -13, -47, -2, -37, -1}, []int{17, 28, 49, -70, -76, -83, -100, -64, 60, 12, -80, -41, -12, 98, 2, 48, 25, 93, 83, 5, 12, -71, 5, 54, 43, 83, 0, -16, -88, -10, 37, 34, -89, -99, -97, -32, 56, 1, -43, 7, -29, 75, 30, 47, -70, -65, -27, 51, -19, 50, 12, -37, 68, 6, 90, 79, 98, -47, -64, 82, 7, 93, -43, -19, 24, 20, 28, 75, 79, 50, 71, 71, 96, 98, -29, 68, -35, 17, 67, 16, 47, -86, 76, 20, -4, 2, -32, -17, 82, -6, 53, 22, -8, -7, 70, -36, -53, 71, 78, 39}, 223308},
}

func TestMaxMultiArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maximumScore(res.Nums, res.Multi)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxMultiArrayNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maximumScore(res.Nums, res.Multi)
		}
	}
}