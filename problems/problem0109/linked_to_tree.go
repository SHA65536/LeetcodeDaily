package problem0109

import (
	. "leetcodedaily/helpers/binarytree"
	. "leetcodedaily/helpers/listnode"
)

/*
Given the head of a singly linked list where elements are sorted in ascending order,
convert it to a height-balanced binary search tree.
*/

func sortedListToBST(head *ListNode) *TreeNode {
	return helper(head, nil)
}

func helper(head, tail *ListNode) *TreeNode {
	var slow, fast *ListNode = head, head
	if head == tail {
		return nil
	}

	// Find the middle
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// Head is the middle
	var res = &TreeNode{Val: slow.Val}
	// Left is left half
	res.Left = helper(head, slow)
	// Right is right half
	res.Right = helper(slow.Next, tail)
	return res
}
