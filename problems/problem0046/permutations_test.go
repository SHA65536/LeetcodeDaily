package problem0046

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected [][]int
}

var Results = []Result{
	{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	{[]int{0, 1}, [][]int{{0, 1}, {1, 0}}},
	{[]int{1}, [][]int{{1}}},
}

func TestPermutations(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := permute(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", got))
	}
}

func BenchmarkPermutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			permute(res.Input)
		}
		permute([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}

func TestPermutationsAlloc(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := permuteAlloc(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPermutationsAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			permuteAlloc(res.Input)
		}
		permuteAlloc([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}
