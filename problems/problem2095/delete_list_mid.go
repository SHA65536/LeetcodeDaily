package problem2095

import (
	. "leetcodedaily/helpers/listnode"
)

/*
You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.
The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing,
where ⌊x⌋ denotes the largest integer less than or equal to x.
*/

func deleteMiddle(head *ListNode) *ListNode {
	var end, mid *ListNode = head, &ListNode{Next: head}
	// Edge case of 1 length list
	if head.Next == nil {
		return nil
	}
	// Moving end each time, and mid every 2nd time
	for i := 0; end.Next != nil; i++ {
		if i%2 == 0 {
			mid = mid.Next
		}
		end = end.Next
	}
	// Removing the node after next
	mid.Next = mid.Next.Next
	return head
}
