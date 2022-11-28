package problem0142

import . "leetcodedaily/helpers/listnode"

/*
Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle.
Note that pos is not passed as a parameter.
Do not modify the linked list.
*/

func detectCycle(head *ListNode) *ListNode {
	var slow, fast, res *ListNode
	slow, fast, res = head, head, head
	if head == nil || head.Next == nil {
		return nil
	}
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next      // Slow goes 1 at a time
		fast = fast.Next.Next // Fast goes 2 at a time
		// If they meet, theres a cycle
		if slow == fast {
			// Finding the cycle
			for slow != res {
				slow = slow.Next
				res = res.Next
			}
			return res
		}
	}
	return nil
}
