package problem1319

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	N        int
	Expected int
}

var Results = []Result{
	{[][]int{{0, 1}, {0, 2}, {1, 2}}, 4, 1},
	{[][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}, 6, 2},
	{[][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}, 6, -1},
}

func TestNumberOfCablesToMove(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := makeConnected(res.N, res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkNumberOfCablesToMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			makeConnected(res.N, res.Input)
		}
	}
}

func TestNumberOfCablesToMoveDFS(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := makeConnectedDFS(res.N, res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkNumberOfCablesToMoveDFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			makeConnectedDFS(res.N, res.Input)
		}
	}
}
