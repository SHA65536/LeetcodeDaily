package problem0019

import . "leetcodedaily/helpers/listnode"

/*
https://leetcode.com/problems/remove-nth-node-from-end-of-list/

Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if removeNthHelper(head, n) == 0 {
		return head.Next
	}
	return head
}

func removeNthHelper(node *ListNode, n int) int {
	if node == nil {
		return n
	}
	cur := removeNthHelper(node.Next, n)
	if cur == 0 {
		node.Next = node.Next.Next
	}
	return cur - 1
}
