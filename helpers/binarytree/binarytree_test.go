package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var treeOne = &TreeNode{Val: 1,
	Left:  &TreeNode{Val: 2},
	Right: &TreeNode{Val: 3}}

var treeTwo = &TreeNode{Val: 1,
	Left: &TreeNode{Val: 2,
		Right: &TreeNode{Val: 3}}}

func TestMakeTree(t *testing.T) {
	assert := assert.New(t)

	res := MakeTree(1, 2, NULL, NULL, 3, NULL, NULL)
	assert.Equal(treeOne, res)
	res = MakeTree(1, 2, NULL, 3, NULL, NULL, NULL)
	assert.Equal(treeTwo, res)
}
