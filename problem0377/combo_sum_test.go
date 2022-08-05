package problem0377

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
	{[]int{1, 2, 3}, 4, 7},
	{[]int{9}, 3, 0},
	{[]int{1, 2, 3, 4, 5}, 10, 464},
	{[]int{1, 2, 3, 4, 5}, 25, 11749641},
}

func TestCombinationSum4(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := combinationSum4(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCombinationSum4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			combinationSum4(res.Input, res.Target)
		}
	}
}

func TestCombinationSum4Naive(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := combinationSum4Naive(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCombinationSum4Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			combinationSum4Naive(res.Input, res.Target)
		}
	}
}
