package problem1305

import . "leetcodedaily/helpers/binarytree"

/*
Given two binary search trees root1 and root2, return a list containing all the integers from both trees sorted in ascending order.
*/

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var arr1, arr2, res []int
	var iter1, iter2 int
	arr1 = make([]int, 0, 5000)
	arr2 = make([]int, 0, 5000)
	if root1 != nil {
		traverseTree(&arr1, root1)
	}
	if root2 != nil {
		traverseTree(&arr2, root2)
	}
	res = make([]int, len(arr1)+len(arr2))
	for iter1+iter2 < (len(arr1) + len(arr2)) {
		if iter1 == len(arr1) {
			res[iter1+iter2] = arr2[iter2]
			iter2++
		} else if iter2 == len(arr2) {
			res[iter1+iter2] = arr1[iter1]
			iter1++
		} else if arr1[iter1] < arr2[iter2] {
			res[iter1+iter2] = arr1[iter1]
			iter1++
		} else {
			res[iter1+iter2] = arr2[iter2]
			iter2++
		}
	}
	return res
}

func traverseTree(res *[]int, root *TreeNode) {
	if root.Left != nil {
		traverseTree(res, root.Left)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		traverseTree(res, root.Right)
	}
}
