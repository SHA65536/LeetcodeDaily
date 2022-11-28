package problem0141

import . "leetcodedaily/helpers/listnode"

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.
*/

const maxVal = 100001

func hasCycle(head *ListNode) bool {
	for head != nil {
		if head.Val == maxVal {
			return true
		}
		head.Val = maxVal
		head = head.Next
	}
	return false
}
