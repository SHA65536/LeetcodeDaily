package problem0304

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Query    [4]int
	Expected int
}

var Results = []Result{
	{[4]int{2, 1, 4, 3}, 8},
	{[4]int{1, 1, 2, 2}, 11},
	{[4]int{1, 2, 2, 4}, 12},
	{[4]int{0, 0, 4, 4}, 58},
}

func TestSumRegion(t *testing.T) {
	assert := assert.New(t)
	var matrix = [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	mat := Constructor(matrix)
	for _, res := range Results {
		want := res.Expected
		got := mat.SumRegion(res.Query[0], res.Query[1], res.Query[2], res.Query[3])
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSumRegion(b *testing.B) {
	var matrix = [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	mat := Constructor(matrix)

	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			mat.SumRegion(res.Query[0], res.Query[1], res.Query[2], res.Query[3])
		}
	}
}

func TestSumRegionPrefix(t *testing.T) {
	assert := assert.New(t)
	var matrix = [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	mat := ConstructorPrefix(matrix)
	for _, res := range Results {
		want := res.Expected
		got := mat.SumRegionPrefix(res.Query[0], res.Query[1], res.Query[2], res.Query[3])
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSumRegionPrefix(b *testing.B) {
	var matrix = [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	mat := ConstructorPrefix(matrix)

	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			mat.SumRegionPrefix(res.Query[0], res.Query[1], res.Query[2], res.Query[3])
		}
	}
}
