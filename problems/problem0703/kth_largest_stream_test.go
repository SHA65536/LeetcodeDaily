package problem0703

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Nums     []int
	K        int
	Stream   []int
	Expected []int
}

var Results = []Result{
	{
		[]int{4, 5, 8, 2}, 3,
		[]int{3, 5, 10, 9, 4},
		[]int{4, 5, 5, 8, 8},
	},
}

func TestKthLargestStream(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]int, len(want))
		kth := Constructor(res.K, res.Nums)
		for i := range res.Stream {
			got[i] = kth.Add(res.Stream[i])
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkKthLargestStream(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			kth := Constructor(res.K, res.Nums)
			for i := range res.Stream {
				kth.Add(res.Stream[i])
			}
		}
	}
}
