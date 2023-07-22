package problem0688

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N, K, Row, Col int
	Expected       float64
}

const resultDetla float64 = 0.00001

var TestCases = []TestCase{
	{3, 2, 0, 0, 0.06250},
	{1, 0, 0, 0, 1},
	{10, 20, 0, 0, 0.00454},
}

func TestKnightOnBoardProbability(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := knightProbability(tc.N, tc.K, tc.Row, tc.Col)
		assert.InDelta(want, got, resultDetla, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkKnightOnBoardProbability(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			knightProbability(tc.N, tc.K, tc.Row, tc.Col)
		}
	}
}
