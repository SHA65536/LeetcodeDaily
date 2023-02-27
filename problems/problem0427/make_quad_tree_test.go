package problem0427

import (
	"fmt"
	. "leetcodedaily/helpers/quadtree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected *Node
}

var Results = []Result{
	{
		Input: [][]int{
			{0, 1},
			{1, 0},
		},
		Expected: &Node{Val: true, IsLeaf: false,
			TopLeft:     &Node{Val: false, IsLeaf: true},
			TopRight:    &Node{Val: true, IsLeaf: true},
			BottomLeft:  &Node{Val: true, IsLeaf: true},
			BottomRight: &Node{Val: false, IsLeaf: true},
		},
	},
	{
		Input: [][]int{
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
		},
		Expected: &Node{
			Val: true, IsLeaf: false,
			TopLeft: &Node{Val: true, IsLeaf: true},
			TopRight: &Node{Val: true, IsLeaf: false,
				TopLeft:     &Node{Val: false, IsLeaf: true},
				TopRight:    &Node{Val: false, IsLeaf: true},
				BottomLeft:  &Node{Val: true, IsLeaf: true},
				BottomRight: &Node{Val: true, IsLeaf: true},
			},
			BottomLeft:  &Node{Val: true, IsLeaf: true},
			BottomRight: &Node{Val: false, IsLeaf: true},
		},
	},
}

func TestMakeQuadTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := construct(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func TestMakeQuadTreeNoGenric(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := constructNoGeneric(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
