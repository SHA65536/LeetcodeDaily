package listnode

import "testing"

func TestListnode(t *testing.T) {
	var original = []int{6, 31, 34, 3, 14, 1, 12, 94, 67, 3, 82, 1, 35, 81, 45, 46, 35, 27, 53, 59, 74, 30, 81, 63, 54, 19, 52, 93, 48, 62, 39, 62, 68, 2, 69, 78, 67, 93, 48, 73, 74, 65, 44, 28, 20, 72, 57, 50, 57, 88, 51, 43, 26, 55, 94, 15, 53, 68, 27, 43, 8, 29, 94, 10, 8, 100, 97, 16, 47, 7, 96, 44, 66, 86, 81, 55, 82, 90, 3, 52, 34, 72, 64, 98, 84, 63, 73, 11, 15, 50, 93, 78, 52, 62, 81, 37, 55, 23, 54, 23}
	var list = MakeListNode(original...)
	for i := range original {
		if original[i] != list.Val {
			t.FailNow()
		}
		list = list.Next
	}
}
