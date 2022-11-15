package problem0222

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
	{MakeTree(1, 2, 4, NULL, NULL, 5, NULL, NULL, 3, 6, NULL, NULL, NULL), 6},
	{MakeTree(), 0},
	{MakeTree(1), 1},
	{MakeTree(1, 2, 4, 8, 16, 32, NULL, NULL, 33, NULL, NULL, 17, 34, NULL, NULL, 35, NULL, NULL, 9, 18, 36, NULL, NULL, 37, NULL, NULL, 19, 38, NULL, NULL, 39, NULL, NULL, 5, 10, 20, 40, NULL, NULL, 41, NULL, NULL, 21, 42, NULL, NULL, 43, NULL, NULL, 11, 22, 44, NULL, NULL, 45, NULL, NULL, 23, 46, NULL, NULL, 47, NULL, NULL, 3, 6, 12, 24, 48, NULL, NULL, 49, NULL, NULL, 25, 50, NULL, NULL, 51, NULL, NULL, 13, 26, 52, NULL, NULL, 53, NULL, NULL, 27, 54, NULL, NULL, 55, NULL, NULL, 7, 14, 28, 56, NULL, NULL, 57, NULL, NULL, 29, 58, NULL, NULL, 59, NULL, NULL, 15, 30, 60, NULL, NULL, 61, NULL, NULL, 31, 62, NULL, NULL, 63, NULL, NULL), 63},
}

func TestCompleteTreeCount(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countNodes(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
func BenchmarkCompleteTreeCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			countNodes(res.Input)
		}
	}
}

func TestCompleteTreeCountNaive(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countNodesNaive(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
func BenchmarkCompleteTreeCountNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			countNodesNaive(res.Input)
		}
	}
}
