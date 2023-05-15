package problem1721

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/listnode"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	K        int
	Expected *ListNode
}

var Results = []Result{
	{MakeListNode(1, 2, 3, 4, 5), 2, MakeListNode(1, 4, 3, 2, 5)},
	{MakeListNode(7, 9, 6, 6, 7, 8, 3, 0, 9, 5), 5, MakeListNode(7, 9, 6, 6, 8, 7, 3, 0, 9, 5)},
	{MakeListNode(1, 2, 3, 4, 5), 4, MakeListNode(1, 4, 3, 2, 5)},
}

func TestSwapNodesLinkedList(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := swapNodes(CopyList(res.Input), res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSwapNodesLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			swapNodes(CopyList(res.Input), res.K)
		}
	}
}
