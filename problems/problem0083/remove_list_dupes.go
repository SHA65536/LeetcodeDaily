package problem0083

import . "leetcodedaily/helpers/listnode"

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for cur := head; cur.Next != nil; {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
