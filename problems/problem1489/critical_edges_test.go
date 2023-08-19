package problem1489

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N        int
	Edges    [][]int
	Expected [][]int
}

var TestCases = []TestCase{
	{5, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}, [][]int{{0, 1}, {2, 3, 4, 5}}},
	{4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}}, [][]int{{}, {0, 1, 2, 3}}},
}

func TestCriticalEdges(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := findCriticalAndPseudoCriticalEdges(tc.N, tc.Edges)
		assert.ElementsMatch(want[0], got[0], fmt.Sprintf("%+v", tc))
		assert.ElementsMatch(want[1], got[1], fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkCriticalEdges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			findCriticalAndPseudoCriticalEdges(tc.N, tc.Edges)
		}
	}
}
