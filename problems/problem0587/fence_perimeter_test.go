package problem0587

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected [][]int
}

var Results = []Result{
	{
		[][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}},
		[][]int{{1, 1}, {2, 0}, {2, 4}, {3, 3}, {4, 2}},
	},
	{
		[][]int{{1, 2}, {2, 2}, {4, 2}},
		[][]int{{1, 2}, {2, 2}, {4, 2}},
	},
}

func TestMakeFence(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := outerTrees(res.Input)
		sort.Slice(got, func(i, j int) bool {
			if got[i][0] == got[j][0] {
				return got[i][1] < got[j][1]
			}
			return got[i][0] < got[j][0]
		})
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
