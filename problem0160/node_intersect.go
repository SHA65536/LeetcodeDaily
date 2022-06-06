package problem0160

import . "leetcodedaily/helpers/listnode"

/*
Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
If the two linked lists have no intersection at all, return null.
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var res, curA, curB *ListNode
	var mapA, mapB = make(map[*ListNode]bool), make(map[*ListNode]bool)
	curA, curB = headA, headB
	for curA != nil || curB != nil {
		if curA != nil {
			mapA[curA] = true
		}
		if curB != nil {
			mapB[curB] = true
		}
		if _, ok := mapA[curB]; ok {
			return curB
		}
		if _, ok := mapB[curA]; ok {
			return curA
		}
		if curA != nil {
			curA = curA.Next
		}
		if curB != nil {
			curB = curB.Next
		}
	}
	return res
}
