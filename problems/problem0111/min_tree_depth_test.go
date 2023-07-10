package problem0111

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    *TreeNode
	Expected int
}

var TestCases = []TestCase{
	{MakeTree(3, 9, NULL, NULL, 20, 15, NULL, NULL, 7, NULL, NULL), 2},
	{MakeTree(2, NULL, 3, NULL, 4, NULL, 5, NULL, 6, NULL, NULL), 5},
}

func TestMinDepthTree(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minDepth(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinDepthTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minDepth(tc.Input)
		}
	}
}
