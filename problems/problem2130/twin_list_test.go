package problem2130

import (
	"fmt"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input *ListNode

	Expected int
}

var Results = []Result{
	{MakeListNode(5, 4, 2, 1), 6},
	{MakeListNode(4, 2, 2, 3), 7},
	{MakeListNode(1, 100000), 100001},
}

func TestMaxTwinLinkedList(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := pairSum(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxTwinLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			pairSum(res.Input)
		}
	}
}
