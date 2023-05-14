package problem1799

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
	{[]int{1, 2}, 1},
	{[]int{3, 4, 6, 8}, 11},
	{[]int{1, 2, 3, 4, 5, 6}, 14},
}

func TestMaxGCDScore(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxScore(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxGCDScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maxScore(res.Input)
		}
	}
}
