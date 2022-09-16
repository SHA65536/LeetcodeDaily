package problem0086

import . "leetcodedaily/helpers/listnode"

/*
https://leetcode.com/problems/partition-list/
Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
You should preserve the original relative order of the nodes in each of the two partitions.
*/

func partition(head *ListNode, x int) *ListNode {
	var first, last *ListNode
	var curf, curl *ListNode
	first, last = &ListNode{}, &ListNode{}
	curf, curl = first, last
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val >= x {
			curl.Next = &ListNode{Val: cur.Val}
			curl = curl.Next
		} else {
			curf.Next = &ListNode{Val: cur.Val}
			curf = curf.Next
		}
	}
	curf.Next = last.Next
	return first.Next
}
