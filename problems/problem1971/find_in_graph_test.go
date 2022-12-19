package problem1971

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N        int
	Edges    [][]int
	Src, Dst int
	Expected bool
}

var Results = []Result{
	{3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2, true},
	{6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5, false},
	{10, [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}, 7, 5, true},
}

func TestFindInGraph(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := validPath(res.N, res.Edges, res.Src, res.Dst)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFindInGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			validPath(res.N, res.Edges, res.Src, res.Dst)
		}
	}
}

func TestFindInGraphQueue(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := validPathQueue(res.N, res.Edges, res.Src, res.Dst)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFindInGraphQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			validPathQueue(res.N, res.Edges, res.Src, res.Dst)
		}
	}
}
