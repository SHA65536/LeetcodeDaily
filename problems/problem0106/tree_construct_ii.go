package problem0106

import . "leetcodedaily/helpers/binarytree"

/*
Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and
postorder is the postorder traversal of the same tree, construct and return the binary tree.
*/

func buildTree(inorder []int, postorder []int) *TreeNode {
	var inorderMap = map[int]int{}
	var construct func(sIn, eIn, sPo, ePo int) *TreeNode

	for i := range inorder {
		inorderMap[inorder[i]] = i
	}

	construct = func(sIn, eIn, sPo, ePo int) *TreeNode {
		if sIn > eIn || sPo > ePo || ePo < 0 {
			return nil
		}
		var root = postorder[ePo]
		var temp = &TreeNode{Val: root}
		var pos = inorderMap[root]

		temp.Left = construct(sIn, pos-1, sPo, sPo+(pos-sIn)-1)
		temp.Right = construct(pos+1, eIn, sPo+(pos-sIn), ePo-1)
		return temp
	}

	return construct(0, len(inorder)-1, 0, len(postorder)-1)
}
