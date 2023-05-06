package problem1498

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Target   int
	Expected int
}

var Results = []Result{
	{[]int{3, 5, 6, 7}, 9, 4},
	{[]int{3, 3, 6, 8}, 10, 6},
	{[]int{2, 3, 3, 4, 6, 7}, 12, 61},
	{make([]int, 74), 10, 320260019},
}

func TestNumSubseqTarget(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numSubseq(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkNumSubseqTarget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			numSubseq(res.Input, res.Target)
		}
	}
}

func TestNumSubseqTargetPrecompute(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numSubseqPrecompute(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkNumSubseqTargetPrecompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			numSubseqPrecompute(res.Input, res.Target)
		}
	}
}
