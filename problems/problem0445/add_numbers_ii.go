package problem0445

import . "leetcodedaily/helpers/listnode"

/*
You are given two non-empty linked lists representing two non-negative integers.
The most significant digit comes first and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{}
	var digits1, digits2 []int
	// Add all digits to our stack
	for cur := l1; cur != nil; cur = cur.Next {
		digits1 = append(digits1, cur.Val)
	}
	for cur := l2; cur != nil; cur = cur.Next {
		digits2 = append(digits2, cur.Val)
	}
	// Loop until no more digits left
	for sum := 0; len(digits1) > 0 || len(digits2) > 0; {
		// If the first list has digits, add to the sum and pop
		if len(digits1) > 0 {
			sum += digits1[len(digits1)-1]
			digits1 = digits1[:len(digits1)-1]
		}
		// If the second list has digits, add to the sum and pop
		if len(digits2) > 0 {
			sum += digits2[len(digits2)-1]
			digits2 = digits2[:len(digits2)-1]
		}
		// Only add the result, don't add the carry
		res.Val = sum % 10
		// Create a new head that points to the current head
		// we set the value to the carry incase this is the last iteration
		res = &ListNode{Val: sum / 10, Next: res}
		// Carry the extra to next round
		sum /= 10
	}
	// Remove leading zero
	if res.Val == 0 {
		return res.Next
	}
	return res
}
