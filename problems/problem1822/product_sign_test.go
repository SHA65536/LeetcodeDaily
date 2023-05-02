package problem1822

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
	{[]int{-1, -2, -3, -4, 3, 2, 1}, 1},
	{[]int{1, 5, 0, 2, -3}, 0},
	{[]int{-1, 1, -1, 1, -1}, -1},
}

func TestProductSign(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := arraySign(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkProductSign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			arraySign(res.Input)
		}
	}
}
