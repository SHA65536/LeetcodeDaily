package problem0429

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/narytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *Node
	Expected [][]int
}

var Results = []Result{
	{&Node{Val: 1, Children: []*Node{
		{Val: 3, Children: []*Node{
			{Val: 5},
			{Val: 6},
		}},
		{Val: 2},
		{Val: 4},
	}}, [][]int{{1}, {3, 2, 4}, {5, 6}}},
}

func TestNaryLevelTraversal(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := levelOrder(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
