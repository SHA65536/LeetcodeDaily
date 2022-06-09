package problem0206

import . "leetcodedaily/helpers/listnode"

/*
https://leetcode.com/problems/reverse-linked-list/

Given the head of a singly linked list, reverse the list, and return the reversed list.
*/

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
