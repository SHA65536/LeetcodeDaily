package problem0523

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Target   int
	Expected bool
}

var Results = []Result{
	{[]int{23, 2, 4, 6, 7}, 6, true},
	{[]int{23, 2, 6, 4, 7}, 6, true},
	{[]int{23, 2, 6, 4, 7}, 13, false},
	{[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 1751423, false},
}

func TestSubarraySumEqual(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := checkSubarraySum(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSubarraySumEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			checkSubarraySum(res.Input, res.Target)
		}
	}
}

func TestSubarraySumEqualOpt(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := checkSubarraySumOpt(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSubarraySumEqualOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			checkSubarraySumOpt(res.Input, res.Target)
		}
	}
}
