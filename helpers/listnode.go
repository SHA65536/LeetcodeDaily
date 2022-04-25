package helpers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// MakeListNode constructs a ListNode from given integers
func MakeListNode(nums ...int) *ListNode {
	resList := &ListNode{}
	cur := resList
	for idx, num := range nums {
		if idx == 0 {
			cur.Val = num
		} else {
			cur.Next = &ListNode{Val: num}
			cur = cur.Next
		}
	}
	return resList
}

func (ln *ListNode) String() string {
	if ln.Next != nil {
		return fmt.Sprintf("%d, %s", ln.Val, ln.Next.String())

	} else {
		return fmt.Sprint(ln.Val)
	}
}
