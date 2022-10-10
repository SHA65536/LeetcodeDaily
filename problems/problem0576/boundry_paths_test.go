package problem0576

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N, M, Max, StartR, StartC int
	Expected                  int
}

var Results = []Result{
	{2, 2, 2, 0, 0, 6},
	{1, 3, 3, 0, 1, 12},
	//{10, 10, 15, 5, 5, 38421910},
}

func TestFindPathsNaive(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findPathsNaive(res.N, res.M, res.Max, res.StartR, res.StartC)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFindPathsNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			findPathsNaive(res.N, res.M, res.Max, res.StartR, res.StartC)
		}
	}
}

func TestFindPathsCache(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findPathsCache(res.N, res.M, res.Max, res.StartR, res.StartC)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFindPathsCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			findPathsCache(res.N, res.M, res.Max, res.StartR, res.StartC)
		}
	}
}
