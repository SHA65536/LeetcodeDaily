package problem0024

import (
	. "leetcodedaily/helpers/listnode"
)

/*
https://leetcode.com/problems/swap-nodes-in-pairs/

Given a linked list, swap every two adjacent nodes and return its head.
You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
*/

func swapPairs(head *ListNode) *ListNode {
	var cur, prev, temp *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	cur = head.Next
	temp = cur.Next
	cur.Next = head
	head.Next = temp
	head = cur
	cur = cur.Next
	for cur.Next != nil && cur.Next.Next != nil {
		prev = cur
		cur = cur.Next
		prev.Next = cur.Next
		temp = cur.Next.Next
		cur.Next.Next = cur
		cur.Next = temp
	}
	return head
}
