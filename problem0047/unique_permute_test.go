package problem0047

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
	{[]int{1, 1, 2},
		[][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	{[]int{1, 2, 3},
		[][]int{{1, 2, 3}, {1, 3, 2}, {2, 3, 1}, {2, 1, 3}, {3, 1, 2}, {3, 2, 1}}},
	{[]int{2, 2, 1, 1},
		[][]int{{2, 2, 1, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {1, 1, 2, 2}, {1, 2, 2, 1}, {1, 2, 1, 2}}},
}

func TestPermuteUnique(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := permuteUnique(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPermuteUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permuteUnique([]int{1, 2, 3, 4, 5, 6, 7, 8})
	}
}
