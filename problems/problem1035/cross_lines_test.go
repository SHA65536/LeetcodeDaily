package problem1035

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	InA, InB []int
	Expected int
}

var Results = []Result{
	{[]int{1, 4, 2}, []int{1, 2, 4}, 2},
	{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
	{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
}

func TestUncrossedLines(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxUncrossedLines(res.InA, res.InB)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkUncrossedLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maxUncrossedLines(res.InA, res.InB)
		}
	}
}
