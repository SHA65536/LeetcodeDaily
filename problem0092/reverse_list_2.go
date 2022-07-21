package problem0092

import . "leetcodedaily/helpers/listnode"

/*
Given the head of a singly linked list and two integers left and right where left <= right,
reverse the nodes of the list from position left to position right, and return the reversed list.
*/

// I went for the 2 pass approach, difference in performance is negligible and its much more readable
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var values []int
	for cur := head; cur != nil; cur = cur.Next {
		values = append(values, cur.Val)
	}
	cur := head
	for idx := range values {
		if idx >= left-1 && idx <= right-1 {
			cur.Val = values[(right-1)-(idx-(left-1))]
		} else {
			cur.Val = values[idx]
		}
		cur = cur.Next
	}
	return head
}
