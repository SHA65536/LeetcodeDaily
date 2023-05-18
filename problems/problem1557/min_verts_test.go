package problem1557

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N        int
	Input    [][]int
	Expected []int
}

var Results = []Result{
	{6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}, []int{0, 3}},
	{5, [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}, []int{0, 2, 3}},
}

func TestMinVertsToReachAll(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findSmallestSetOfVertices(res.N, res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMinVertsToReachAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			findSmallestSetOfVertices(res.N, res.Input)
		}
	}
}
