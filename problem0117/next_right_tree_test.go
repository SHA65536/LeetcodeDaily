package problem0117

import (
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	assert := assert.New(t)

	var input1 = &Node{Val: 1,
		Left: &Node{Val: 2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5}},
		Right: &Node{Val: 3,
			Right: &Node{Val: 7}}}
	var expected1 = &Node{Val: 1,
		Left: &Node{Val: 2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5}},
		Right: &Node{Val: 3,
			Right: &Node{Val: 7}}}
	expected1.Left.Next = expected1.Right
	expected1.Left.Left.Next = expected1.Left.Right
	expected1.Left.Right.Next = expected1.Right.Right

	got := connect(input1)
	assert.Equal(expected1, got)
}
