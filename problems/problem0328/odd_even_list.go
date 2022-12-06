package problem0328

import . "leetcodedaily/helpers/listnode"

/*
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.
The first node is considered odd, and the second node is even, and so on.
Note that the relative order inside both the even and odd groups should remain as it was in the input.
You must solve the problem in O(1) extra space complexity and O(n) time complexity.
*/

func oddEvenList(head *ListNode) *ListNode {
	odd := &ListNode{}
	headOdd := odd
	even := &ListNode{}
	headEven := even
	for i := 0; head != nil; i++ {
		if i%2 == 0 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
		head = head.Next
	}
	even.Next = nil
	odd.Next = headEven.Next
	return headOdd.Next
}
