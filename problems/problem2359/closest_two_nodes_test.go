package problem2359

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Edges    []int
	E1, E2   int
	Expected int
}

var Results = []Result{
	{[]int{2, 2, 3, -1}, 0, 1, 2},
	{[]int{1, 2, -1}, 0, 2, 2},
	{[]int{5, 4, 5, 4, 3, 6, -1}, 0, 1, -1},
	{[]int{4, 4, 8, -1, 9, 8, 4, 4, 1, 1}, 5, 6, 1},
}

func TestClosestNodeToTwoNodes(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := closestMeetingNode(res.Edges, res.E1, res.E2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkClosestNodeToTwoNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			closestMeetingNode(res.Edges, res.E1, res.E2)
		}
	}
}
