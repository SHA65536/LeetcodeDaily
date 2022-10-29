package problem2136

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Plant    []int
	Grow     []int
	Expected int
}

var Results = []Result{
	{[]int{1, 4, 3}, []int{2, 3, 1}, 9},
	{[]int{1, 2, 3, 2}, []int{2, 1, 2, 1}, 9},
	{[]int{1}, []int{1}, 2},
	{[]int{4, 1, 1, 10, 5, 5, 5, 9, 5, 1}, []int{4, 1, 4, 7, 10, 7, 6, 5, 6, 7}, 49},
}

func TestEarliestBloom(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := earliestFullBloom(res.Plant, res.Grow)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
