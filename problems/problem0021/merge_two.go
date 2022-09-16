package problem0021

import (
	. "leetcodedaily/helpers/listnode"
)

/*
https://leetcode.com/problems/merge-two-sorted-lists/

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list.
The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, temp *ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	temp = head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			temp.Next = list2
			temp, list2 = temp.Next, list2.Next
		} else {
			temp.Next = list1
			temp, list1 = temp.Next, list1.Next
		}
	}
	if list1 != nil {
		temp.Next = list1
	} else {
		temp.Next = list2
	}
	return head
}
