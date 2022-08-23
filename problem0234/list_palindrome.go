package problem0234

import . "leetcodedaily/helpers/listnode"

func isPalindrome(head *ListNode) bool {
	var copy **ListNode = &head
	return checkPalindrome(copy, head)
}

func checkPalindrome(head **ListNode, temp *ListNode) bool {
	if temp == nil {
		return true
	}
	if checkPalindrome(head, temp.Next) {
		if (*head).Val == temp.Val {
			*head = (*head).Next
			return true
		}
	}
	return false
}
