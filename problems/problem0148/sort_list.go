package problem0148

import (
	. "leetcodedaily/helpers/listnode"
	"sort"
)

/*
Given the head of a linked list, return the list after sorting it in ascending order.
*/

func sortList(head *ListNode) *ListNode {
	var res []int
	var temp = head
	for temp != nil {
		res = append(res, temp.Val)
		temp = temp.Next
	}
	sort.Ints(res)
	temp = head
	for i := range res {
		temp.Val = res[i]
		temp = temp.Next
	}
	return head
}

func sortListInPlace(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Splitting the list into two using slow and fast approach
	var temp, slow, fast *ListNode = nil, head, head
	for fast != nil && fast.Next != nil {
		temp = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	temp.Next = nil

	// Sorting each half
	l1, l2 := sortListInPlace(head), sortListInPlace(slow)

	// Merging them
	return mergeList(l1, l2)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	var ptr = &ListNode{}
	var cur = ptr

	// Looping over the lists, and merging them into ptr
	// in ascending order
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	// If lists aren't equal size, add the rest
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return ptr.Next
}
