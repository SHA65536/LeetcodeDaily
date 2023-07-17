package problem0445

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/listnode"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	L1, L2   *ListNode
	Expected *ListNode
}

var TestCases = []TestCase{
	{
		MakeListNode(7, 2, 4, 3),
		MakeListNode(5, 6, 4),
		MakeListNode(7, 8, 0, 7),
	},
	{
		MakeListNode(2, 4, 3),
		MakeListNode(5, 6, 4),
		MakeListNode(8, 0, 7),
	},
	{
		MakeListNode(0),
		MakeListNode(0),
		MakeListNode(0),
	},
	{
		MakeListNode(2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9),
		MakeListNode(5, 6, 4, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9, 9, 9, 9),
		MakeListNode(8, 0, 7, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 7, 2, 4, 3, 8),
	},
	{
		MakeListNode(5, 0, 0, 0),
		MakeListNode(5, 0, 0, 0),
		MakeListNode(1, 0, 0, 0, 0),
	},
}

func TestAddTwoNumbersII(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := addTwoNumbers(tc.L1, tc.L2)
		assert.Equal(want.String(), got.String(), fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkAddTwoNumbersII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			addTwoNumbers(tc.L1, tc.L2)
		}
	}
}
