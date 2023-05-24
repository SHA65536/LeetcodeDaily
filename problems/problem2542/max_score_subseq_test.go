package problem2542

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Nums1, Nums2 []int
	K            int
	Expected     int64
}

var Results = []Result{
	{[]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3, 12},
	{[]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1, 30},
}

func TestMaxScoreSubsequence(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxScore(res.Nums1, res.Nums2, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxScoreSubsequenceh(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maxScore(res.Nums1, res.Nums2, res.K)
		}
	}
}
