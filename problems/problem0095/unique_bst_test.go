package problem0095

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    int
	Expected []*TreeNode
}

var TestCases = []TestCase{
	{1, []*TreeNode{MakeTree(1)}},
	{3, []*TreeNode{
		MakeTree(1, NULL, 3, 2, NULL, NULL, NULL),
		MakeTree(1, NULL, 2, NULL, 3, NULL, NULL),
		MakeTree(2, 1, NULL, NULL, 3, NULL, NULL),
		MakeTree(3, 2, 1, NULL, NULL, NULL, NULL),
		MakeTree(3, 1, NULL, 2, NULL, NULL, NULL),
	}},
}

func TestUniqueBSTs(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := generateTrees(tc.Input)
		assert.ElementsMatch(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkUniqueBSTs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			generateTrees(tc.Input)
		}
	}
}
