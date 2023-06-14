package problem0530

import (
	. "leetcodedaily/helpers/binarytree"
	"math"
)

/*
Given the root of a Binary Search Tree (BST),
return the minimum absolute difference between the values of any two different nodes in the tree.
*/

func getMinimumDifferencePointers(root *TreeNode) int {
	var res, prev int = math.MaxInt, math.MaxInt
	helperWithPointers(root, &prev, &res)
	return res
}

func helperWithPointers(root *TreeNode, prev, min *int) {
	if root == nil {
		return
	}
	helperWithPointers(root.Left, prev, min)
	*min = minWithAbs((root.Val - *prev), *min)
	*prev = root.Val
	helperWithPointers(root.Right, prev, min)
}

func getMinimumDifferenceChan(root *TreeNode) int {
	var res, prev int = math.MaxInt, math.MaxInt
	var ch = make(chan int)
	go func() {
		helperWithChans(root, ch)
		close(ch)
	}()
	for cur := range ch {
		res = minWithAbs(prev-cur, res)
		prev = cur
	}
	return res
}

func helperWithChans(root *TreeNode, ch chan int) {
	if root != nil {
		helperWithChans(root.Left, ch)
		ch <- root.Val
		helperWithChans(root.Right, ch)
	}
}

func minWithAbs(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a < b {
		return a
	}
	return b
}
