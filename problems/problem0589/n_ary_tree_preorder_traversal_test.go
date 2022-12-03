package problem0589

import (
	"fmt"
	. "leetcodedaily/helpers/narytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *Node
	Expected []int
}

var Results = []Result{
	{&Node{Val: 1, Children: []*Node{
		{Val: 3, Children: []*Node{
			{Val: 5},
			{Val: 6},
		}},
		{Val: 2},
		{Val: 4},
	}}, []int{1, 3, 5, 6, 2, 4}},
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
	}},
		[]int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10}},
}

func TestPreOrderNary(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := preorder(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPreOrderNary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			preorder(res.Input)
		}
	}
}

func TestPreOrderNaryShared(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := preorderShared(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPreOrderNaryShared(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			preorderShared(res.Input)
		}
	}
}

func TestPreOrderNaryStack(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := preorderStack(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPreOrderNaryStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			preorderStack(res.Input)
		}
	}
}
