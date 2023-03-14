package problem0129

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
	{MakeTree(1, 2, NULL, NULL, 3, NULL, NULL), 25},
	{MakeTree(4, 9, 5, NULL, NULL, 1, NULL, NULL, 0, NULL, NULL), 1026},
}

func TestSumRootToLeafPaths(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		input := CopyTree(res.Input)
		got := sumNumbers(input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSumRootToLeafPaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			input := CopyTree(res.Input)
			sumNumbers(input)
		}
	}
}

func TestSumRootToLeafPathsIter(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		input := CopyTree(res.Input)
		got := sumNumbersIterative(input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSumRootToLeafPathsIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			input := CopyTree(res.Input)
			sumNumbersIterative(input)
		}
	}
}
