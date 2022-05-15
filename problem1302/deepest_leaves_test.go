package problem1302

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Root     *TreeNode
	Expected int
}

var Results = []Result{
	{&TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 7}},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3,
			Right: &TreeNode{Val: 6,
				Right: &TreeNode{Val: 8}}}}, 15},
	{&TreeNode{Val: 6,
		Left: &TreeNode{Val: 7,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 9}},
			Right: &TreeNode{Val: 7,
				Left: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8,
			Left: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 1}},
			Right: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 4}}}}, 19},
}

func TestDeepestLeavesSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := deepestLeavesSum(res.Root)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkDeepestLeavesSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deepestLeavesSum(Results[0].Root)
		deepestLeavesSum(Results[1].Root)
	}
}

func BenchmarkDeepestLeavesSumTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deepestLeavesSumTwo(Results[0].Root)
		deepestLeavesSumTwo(Results[1].Root)
	}
}

/*
Constraints:
The number of nodes in the tree is in the range [1, 104].
1 <= Node.val <= 100
*/
