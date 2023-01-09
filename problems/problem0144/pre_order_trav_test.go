package problem0144

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected []int
}

var Results = []Result{
	{MakeTree(1, NULL, 2, 3, NULL, NULL, NULL), []int{1, 2, 3}},
	{MakeTree(), []int{}},
	{MakeTree(1), []int{1}},
	{MakeTree(25, 15, 10, 4, NULL, NULL, 12, NULL, NULL, 22, 18, NULL, NULL, 24, NULL, NULL, 50, 35, 31, NULL, NULL, 44, NULL, NULL, 70, 66, NULL, NULL, 90, NULL, NULL),
		[]int{25, 15, 10, 4, 12, 22, 18, 24, 50, 35, 31, 44, 70, 66, 90}},
}

func TestPreOrderTraversal(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := preorderTraversal(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPreOrderTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			preorderTraversal(res.Input)
		}
	}
}

func TestPreOrderTraversalStack(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := preorderTraversalStack(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPreOrderTraversalStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			preorderTraversalStack(res.Input)
		}
	}
}
