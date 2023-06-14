package problem0559

import (
	"fmt"
	. "leetcodedaily/helpers/narytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    *Node
	Expected int
}

var TestCases = []TestCase{
	{&Node{Val: 1, Children: []*Node{
		{Val: 3, Children: []*Node{
			{Val: 5},
			{Val: 6},
		}},
		{Val: 2},
		{Val: 4},
	}}, 3},
	{&Node{Val: 1, Children: []*Node{
		{Val: 2},
		{Val: 3, Children: []*Node{
			{Val: 6},
			{Val: 7, Children: []*Node{
				{Val: 11, Children: []*Node{
					{Val: 14},
				}},
			}},
		}},
		{Val: 4, Children: []*Node{
			{Val: 8, Children: []*Node{
				{Val: 12},
			}},
		}},
		{Val: 5, Children: []*Node{
			{Val: 9, Children: []*Node{
				{Val: 13},
			}},
			{Val: 10},
		}},
	}}, 5},
}

func TestMaxDepthOfNaryTree(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maxDepth(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxDepthOfNaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maxDepth(tc.Input)
		}
	}
}
