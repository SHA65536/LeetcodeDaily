package problem0783

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected int
}

var Results = []Result{
	{MakeTree(4, 2, 1, NULL, NULL, 3, NULL, NULL, 6, NULL, NULL), 1},
	{MakeTree(1, 0, NULL, NULL, 48, 12, NULL, NULL, 49, NULL, NULL), 1},
	{MakeTree(5, 1, NULL, NULL, 48, 12, NULL, NULL, 129, NULL, NULL), 4},
}

func TestMinDistanceInBST(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minDiffInBST(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMinDistanceInBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			minDiffInBST(res.Input)
		}
	}
}

func TestMinDistanceInBSTStack(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minDiffInBSTStack(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMinDistanceInBSTStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			minDiffInBSTStack(res.Input)
		}
	}
}
