package problem0025

import . "leetcodedaily/helpers/listnode"

/*
Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
k is a positive integer and is less than or equal to the length of the linked list.
If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
You may not alter the values in the list's nodes, only nodes themselves may be changed.
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	var cur, prev *ListNode
	var count int
	for cur = head; count < k; count++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	for prev = reverseKGroup(cur, k); count > 0; count-- {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
