package problem1579

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N        int
	Edges    [][]int
	Expected int
}

var Results = []Result{
	{4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}, 2},
	{4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}}, 0},
	{4, [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}}, -1},
}

func TestMaxEdgesRemovedTraversal(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxNumEdgesToRemove(res.N, copyIntInt(res.Edges))
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxEdgesRemovedTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maxNumEdgesToRemove(res.N, copyIntInt(res.Edges))
		}
	}
}

func copyIntInt(a [][]int) [][]int {
	res := make([][]int, len(a))
	for i := range a {
		res[i] = make([]int, len(a[i]))
		copy(res[i], a[i])
	}
	return res
}
