package problem1514

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N          int
	Edges      [][]int
	Prob       []float64
	Start, End int
	Expected   float64
}

var TestCases = []TestCase{
	{
		N:     3,
		Edges: [][]int{{0, 1}, {1, 2}, {0, 2}},
		Prob:  []float64{0.5, 0.5, 0.2},
		Start: 0, End: 2,
		Expected: 0.25,
	},
	{
		N:     3,
		Edges: [][]int{{0, 1}, {1, 2}, {0, 2}},
		Prob:  []float64{0.5, 0.5, 0.3},
		Start: 0, End: 2,
		Expected: 0.3,
	},
	{
		N:     3,
		Edges: [][]int{{0, 1}},
		Prob:  []float64{0.5},
		Start: 0, End: 2,
		Expected: 0,
	},
}

const precision float64 = 0.00001

func TestMaxProbabilityPath(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := maxProbability(tc.N, tc.Edges, tc.Prob, tc.Start, tc.End)
		assert.InDelta(want, got, precision, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMaxProbabilityPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			maxProbability(tc.N, tc.Edges, tc.Prob, tc.Start, tc.End)
		}
	}
}
