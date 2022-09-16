package problem0876

import . "leetcodedaily/helpers/listnode"

/*
https://leetcode.com/problems/middle-of-the-linked-list/

Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.
*/

func middleNode(head *ListNode) *ListNode {
	_, res := getMiddle(head, 0)
	return res
}

func getMiddle(node *ListNode, cur int) (int, *ListNode) {
	if node == nil {
		return 0, nil
	}
	val, res := getMiddle(node.Next, cur+1)
	// After middle is found we just return mid
	if val == cur || val+1 == cur {
		return -2, node
	}
	if res != nil {
		return -2, res
	}
	return val + 1, nil
}
