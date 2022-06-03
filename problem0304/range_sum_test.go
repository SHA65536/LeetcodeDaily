package problem0304

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Query    [4]int
	Expected int
}

var Matrix = [][]int{
	{3, 0, 1, 4, 2},
	{5, 6, 3, 2, 1},
	{1, 2, 0, 1, 5},
	{4, 1, 0, 1, 7},
	{1, 0, 3, 0, 5},
}

var Results = []Result{
	{[4]int{2, 1, 4, 3}, 8},
	{[4]int{1, 1, 2, 2}, 11},
	{[4]int{1, 2, 2, 4}, 12},
	{[4]int{0, 0, 4, 4}, 58},
}

func TestTranspose(t *testing.T) {
	assert := assert.New(t)

	mat := Constructor(Matrix)
	for _, res := range Results {
		want := res.Expected
		got := mat.SumRegion(res.Query[0], res.Query[1], res.Query[2], res.Query[3])
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
