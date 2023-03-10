package problem0382

import (
	. "leetcodedaily/helpers/listnode"
	"math/rand"
)

type Solution struct {
	Head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{Head: head}
}

func (s *Solution) GetRandom() int {
	res, n := 0, 0
	for h := s.Head; h != nil; h = h.Next {
		n++
		if rand.Intn(n) == 0 {
			res = h.Val
		}
	}
	return res
}
