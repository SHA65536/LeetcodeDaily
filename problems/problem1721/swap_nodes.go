package problem1721

import . "leetcodedaily/helpers/listnode"

/*
You are given the head of a linked list, and an integer k.
Return the head of the linked list after swapping the values of the kth node from the
beginning and the kth node from the end (the list is 1-indexed).
*/

func swapNodes(head *ListNode, k int) *ListNode {
	var cur, source, target *ListNode
	cur = head
	// Finding the source
	for i := 1; cur.Next != nil; i++ {
		if i == k {
			break
		}
		cur = cur.Next
	}
	source = cur
	// Finding the target
	for target = head; cur.Next != nil; {
		cur = cur.Next
		target = target.Next
	}
	// Swapping
	source.Val, target.Val = target.Val, source.Val
	return head
}
