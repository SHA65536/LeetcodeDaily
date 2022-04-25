package add_two_numbers

import (
	"fmt"
	. "leetcodedaily/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AddTwoNumbersResult struct {
	L1       *ListNode
	L2       *ListNode
	Expected *ListNode
}

var AddTwoNumbersResults = []AddTwoNumbersResult{
	{
		L1:       MakeListNode(2, 4, 3),
		L2:       MakeListNode(5, 6, 4),
		Expected: MakeListNode(7, 0, 8),
	},
	{
		L1:       MakeListNode(0),
		L2:       MakeListNode(0),
		Expected: MakeListNode(0),
	},
	{
		L1:       MakeListNode(1),
		L2:       MakeListNode(1),
		Expected: MakeListNode(2),
	},
	{
		L1:       MakeListNode(9, 9, 9, 9, 9, 9, 9),
		L2:       MakeListNode(9, 9, 9, 9),
		Expected: MakeListNode(8, 9, 9, 9, 0, 0, 0, 1),
	},
	{
		L1:       MakeListNode(9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9),
		L2:       MakeListNode(9, 9, 9, 9),
		Expected: MakeListNode(8, 9, 9, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1),
	},
}

func TestAddTwoNumbers(t *testing.T) {
	assert := assert.New(t)
	addTwoPrinter := func(r AddTwoNumbersResult, g *ListNode) string {
		return fmt.Sprintf(
			"L1 - %s\nL2 - %s\nExp - %s\nGot - %s",
			r.L1.String(),
			r.L2.String(),
			r.Expected.String(),
			g.String())
	}

	for _, res := range AddTwoNumbersResults {
		want := res.Expected
		got := addTwoNumbers(res.L1, res.L2)
		assert.Equal(want, got, addTwoPrinter(res, got))
	}
}
