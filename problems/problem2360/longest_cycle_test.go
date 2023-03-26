package problem2360

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected int
}

var Results = []Result{
	{[]int{3, 3, 4, 2, 3}, 3},
	{[]int{2, -1, 3, 1}, -1},
	{[]int{2, 3, 3, 1}, 2},
}

func TestLongestCycle(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := longestCycle(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkLongestCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			longestCycle(res.Input)
		}
	}
}
