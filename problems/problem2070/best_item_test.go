package problem2070

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Items    [][]int
	Queries  []int
	Expected []int
}

var Results = []Result{
	{[][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 5, 5, 6, 6}},
	{[][]int{{1, 2}, {1, 2}, {1, 3}, {1, 4}}, []int{1}, []int{4}},
	{[][]int{{10, 1000}}, []int{5}, []int{0}},
	{[][]int{{1, 2}, {1, 2}, {1, 3}, {1, 4}, {1, 5}}, []int{1, 2, 3}, []int{5, 5, 5}},
}

func TestMostBeautifulItem(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := maximumBeauty(res.Items, res.Queries)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
