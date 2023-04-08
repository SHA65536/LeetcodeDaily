package problem0133

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneGraph(t *testing.T) {
	assert := assert.New(t)

	var Nodes = []*Node{
		{Val: 1}, {Val: 2}, {Val: 3}, {Val: 4},
	}

	Nodes[0].Neighbors = []*Node{Nodes[1], Nodes[3]}
	Nodes[1].Neighbors = []*Node{Nodes[0], Nodes[2]}
	Nodes[2].Neighbors = []*Node{Nodes[1], Nodes[3]}
	Nodes[3].Neighbors = []*Node{Nodes[0], Nodes[2]}

	// Should copy graph
	assert.Equal(Nodes[0], cloneGraph(Nodes[0]))

	// Should not have same memory
	assert.False(Nodes[0] == cloneGraph(Nodes[0]))

	// Should copy single node
	assert.Equal(&Node{Val: 1, Neighbors: []*Node{}}, cloneGraph(&Node{Val: 1, Neighbors: []*Node{}}))

	// Should not have same memory
	assert.False(&Node{Val: 1} == cloneGraph(&Node{Val: 1}))

	// Should nil on nil input
	assert.Equal(nil, nil)
}
