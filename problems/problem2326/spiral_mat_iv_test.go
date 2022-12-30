package problem2326

import (
	"fmt"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Rows, Cols int
	List       *ListNode
	Expected   [][]int
}

var Results = []Result{
	{3, 5, MakeListNode(3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0), [][]int{
		{3, 0, 2, 6, 8},
		{5, 0, -1, -1, 1},
		{5, 2, 4, 9, 7},
	}},
	{1, 4, MakeListNode(0, 1, 2), [][]int{{0, 1, 2, -1}}},
}

func TestSpiralMatrixIV(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := spiralMatrix(res.Rows, res.Cols, res.List)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
