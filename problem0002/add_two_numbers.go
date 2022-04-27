package problem0002

import . "leetcodedaily/helpers"

/*
https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	newList := &ListNode{}
	cur := newList
	for l1 != nil || l2 != nil {
		var aVal, bVal, rem int
		// getting digits from current node
		// if node does not exist, digit is 0
		if l1 != nil {
			aVal = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			bVal = l2.Val
			l2 = l2.Next
		}

		// summing up the digits
		rem = (aVal + bVal + carry) % 10
		// carry goes into next node
		carry = (aVal + bVal + carry) / 10

		cur.Next = &ListNode{Val: rem}
		cur = cur.Next
	}
	// accounting for last digits summing up to more than 10
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return newList.Next
}
