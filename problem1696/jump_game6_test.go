package problem1696

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	K        int
	Expected int
}

var Results = []Result{
	{[]int{1, -1, -2, 4, -7, 3}, 2, 7},
	{[]int{10, -5, -2, 4, 0, 3}, 3, 17},
	{[]int{1, -5, -20, 4, -1, 3, -6, -3}, 2, 0},
	{[]int{4, 11, 15, 5, 2, 13, 4, 8, 14, 18, 18, 20, -3, -9, -6, 19, 2, -9, -14, 20, 6, -1, -19, -7, 0, -13, -15, -12, -12, -14, 20, 3, 4, -12, 16, -10, -7, -6, -13, 11}, 8, 233},
	{[]int{1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, 15, -4},
}

func TestMaxResult(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxResult(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxResult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maxResult(res.Input, res.K)
		}
	}
}

func TestMaxResultOptimized(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxResultOptimized(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxResultOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maxResultOptimized(res.Input, res.K)
		}
	}
}

func TestMaxResultTabulation(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxResultTabulation(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxResultTabulation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maxResultTabulation(res.Input, res.K)
		}
	}
}

func TestMaxResultDoubleQueue(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxResultDoubleQueue(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxResultDoubleQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maxResultDoubleQueue(res.Input, res.K)
		}
	}
}
