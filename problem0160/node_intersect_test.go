package problem0160

import (
	"fmt"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	L1       *ListNode
	L2       *ListNode
	Expected *ListNode
}

var Results = []Result{
	{MakeListNode(1, 2, 3), MakeListNode(1, 2, 3), MakeListNode(4, 5, 6)},
	{MakeListNode(1, 2), MakeListNode(1, 2, 3), MakeListNode(4, 5, 6)},
	{MakeListNode(1, 2, 3), MakeListNode(1, 2), MakeListNode(4, 5, 6)},
}

func TestGetIntersectionNode(t *testing.T) {
	assert := assert.New(t)
	Results[0].L1.Next.Next.Next = Results[0].Expected
	Results[0].L2.Next.Next.Next = Results[0].Expected

	Results[1].L1.Next.Next = Results[1].Expected
	Results[1].L2.Next.Next.Next = Results[1].Expected

	Results[2].L1.Next.Next.Next = Results[2].Expected
	Results[2].L2.Next.Next = Results[2].Expected

	for _, res := range Results {
		want := res.Expected
		got := getIntersectionNode(res.L1, res.L2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
