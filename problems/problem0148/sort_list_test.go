package problem0148

import (
	"fmt"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	Expected *ListNode
}

var Results = []Result{
	{MakeListNode(4, 2, 1, 3), MakeListNode(1, 2, 3, 4)},
	{MakeListNode(-1, 5, 3, 4, 0), MakeListNode(-1, 0, 3, 4, 5)},
}

func TestSortList(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := CopyList(res.Expected)
		got := sortList(CopyList(res.Input))
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSortList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			sortList(CopyList(res.Input))
		}
	}
}

func TestSortListInPlace(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := CopyList(res.Expected)
		got := sortListInPlace(CopyList(res.Input))
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSortListInPlace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			sortListInPlace(CopyList(res.Input))
		}
	}
}
