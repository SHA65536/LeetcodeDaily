package listnode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// MakeListNode constructs a ListNode from given integers
func MakeListNode(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
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

// CopyList creates a deep copy of the given list
func CopyList(head *ListNode) *ListNode {
	var res *ListNode = &ListNode{}
	var temp = res
	for head != nil {
		temp.Next = &ListNode{Val: head.Val}
		temp = temp.Next
		head = head.Next
	}
	return res.Next
}

func (ln *ListNode) String() string {
	if ln == nil {
		return ""
	}
	if ln.Next != nil {
		return fmt.Sprintf("%d, %s", ln.Val, ln.Next.String())
	}
	return fmt.Sprint(ln.Val)

}
