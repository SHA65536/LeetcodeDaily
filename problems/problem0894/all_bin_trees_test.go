package problem0894

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N        int
	Expected []*TreeNode
}

var TestCases = []TestCase{
	{
		7, []*TreeNode{
			MakeTree(0, 0, NULL, NULL, 0, 0, NULL, NULL, 0, 0, NULL, NULL, 0, NULL, NULL),
			MakeTree(0, 0, NULL, NULL, 0, 0, 0, NULL, NULL, 0, NULL, NULL, 0, NULL, NULL),
			MakeTree(0, 0, 0, NULL, NULL, 0, NULL, NULL, 0, 0, NULL, NULL, 0, NULL, NULL),
			MakeTree(0, 0, 0, NULL, NULL, 0, 0, NULL, NULL, 0, NULL, NULL, 0, NULL, NULL),
			MakeTree(0, 0, 0, 0, NULL, NULL, 0, NULL, NULL, 0, NULL, NULL, 0, NULL, NULL),
		},
	},
	{
		3, []*TreeNode{
			MakeTree(0, 0, NULL, NULL, 0, NULL, NULL),
		},
	},
}

func TestAllFullBinaryTrees(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := allPossibleFBT(tc.N)
		assert.ElementsMatch(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkAllFullBinaryTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			allPossibleFBT(tc.N)
		}
	}
}
