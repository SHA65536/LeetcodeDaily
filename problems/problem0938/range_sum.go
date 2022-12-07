package problem0938

import . "leetcodedaily/helpers/binarytree"

func rangeSumBST(root *TreeNode, low int, high int) int {
	var sum int
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}
	if root.Val > low && root.Left != nil {
		sum += rangeSumBST(root.Left, low, high)
	}
	if root.Val < high && root.Right != nil {
		sum += rangeSumBST(root.Right, low, high)
	}
	return sum
}
