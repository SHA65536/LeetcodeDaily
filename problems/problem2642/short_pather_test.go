package problem2642

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Commands []string
	N        int
	Edges    [][]int
	Values   [][]int
	Expected []int
}

var TestCases = []TestCase{
	{
		Commands: []string{"Graph", "shortestPath", "shortestPath", "addEdge", "shortestPath"},
		N:        4,
		Edges:    [][]int{{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}},
		Values:   [][]int{{}, {3, 2}, {0, 3}, {1, 3, 4}, {0, 3}},
		Expected: []int{0, 6, -1, 0, 6},
	},
}

func TestShortPathGraph(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var want = tc.Expected
		var got = make([]int, len(want))
		var graph = Constructor(tc.N, tc.Edges)
		for i, cmd := range tc.Commands {
			switch cmd {
			case "shortestPath":
				got[i] = graph.ShortestPath(tc.Values[i][0], tc.Values[i][1])
			case "addEdge":
				graph.AddEdge(tc.Values[i])
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkShortPathGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			var graph = Constructor(tc.N, tc.Edges)
			for i, cmd := range tc.Commands {
				switch cmd {
				case "shortestPath":
					graph.ShortestPath(tc.Values[i][0], tc.Values[i][1])
				case "addEdge":
					graph.AddEdge(tc.Values[i])
				}
			}
		}
	}
}
