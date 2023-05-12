package problem2140

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected int64
}

var Results = []Result{
	{[][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}, 5},
	{[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, 7},
}

func TestMaxQuizPoints(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := mostPoints(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxQuizPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			mostPoints(res.Input)
		}
	}
}

func TestMaxQuizPointsIterative(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := mostPointsIterative(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxQuizPointsIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			mostPointsIterative(res.Input)
		}
	}
}
