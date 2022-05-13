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

func (ln *ListNode) String() string {
	if ln == nil {
		return ""
	}
	if ln.Next != nil {
		return fmt.Sprintf("%d, %s", ln.Val, ln.Next.String())
	}
	return fmt.Sprint(ln.Val)

}
