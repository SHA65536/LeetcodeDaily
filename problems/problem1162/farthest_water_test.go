package problem1162

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected int
}

var Results = []Result{
	{[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, 2},
	{[][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 4},
	{[][]int{{0}}, -1},
	{[][]int{{1}}, -1},
	{func() [][]int {
		res := make([][]int, 100)
		for i := range res {
			res[i] = make([]int, 100)
		}
		res[0][0] = 1
		return res
	}(), 198},
}

func TestFarthestWaterCell(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := maxDistance(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
