package problem2017

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected int64
}

var Results = []Result{
	{[][]int{{2, 5, 4}, {1, 5, 1}}, 4},
	{[][]int{{3, 3, 1}, {8, 5, 2}}, 4},
	{[][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}, 7},
}

func TestGridGame(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := gridGame(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
