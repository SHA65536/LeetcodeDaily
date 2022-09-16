package problem0023

import (
	. "leetcodedaily/helpers/listnode"
)

/*
https://leetcode.com/problems/merge-k-sorted-lists/

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.
*/

func mergeKLists(lists []*ListNode) *ListNode {
	var head = &ListNode{}
	var cur = head
	var min, minIdx int
	var stale bool
	for !stale {
		stale = true
		for idx, node := range lists {
			if node != nil {
				if stale {
					min = node.Val
					minIdx = idx
					stale = false
				} else if min > node.Val {
					min = node.Val
					minIdx = idx
				}
			}
		}
		if !stale {
			cur.Next = &ListNode{Val: min}
			cur = cur.Next
			if lists[minIdx].Next == nil {
				lists[minIdx] = nil
			} else {
				*lists[minIdx] = *lists[minIdx].Next
			}
		}
	}
	return head.Next
}
