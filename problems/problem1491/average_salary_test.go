package problem1491

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected float64
}

var Results = []Result{
	{[]int{4000, 3000, 1000, 2000}, 2500},
	{[]int{1000, 2000, 3000}, 2000},
}

func TestMaxEdgesRemovedTraversal(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := average(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxEdgesRemovedTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			average(res.Input)
		}
	}
}
