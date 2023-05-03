package problem2215

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	In1, In2 []int
	Expected [][]int
}

var Results = []Result{
	{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
	{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, nil}},
}

func TestArrayDifference(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findDifference(res.In1, res.In2)
		sort.Ints(got[0])
		sort.Ints(got[1])
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkArrayDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			findDifference(res.In1, res.In2)
		}
	}
}
