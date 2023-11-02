package problem2265

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    *TreeNode
	Expected int
}

var TestCases = []TestCase{
	{MakeTree(4, 8, 0, NULL, NULL, 1, NULL, NULL, 5, NULL, 6, NULL, NULL), 5},
	{MakeTree(1, NULL, NULL), 1},
}

func TestTreeNodeAvg(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := averageOfSubtree(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkTreeNodeAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			averageOfSubtree(tc.Input)
		}
	}
}
