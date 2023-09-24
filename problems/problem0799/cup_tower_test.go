package problem0799

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Poured, QRow, QGlass int
	Expected             float64
}

const delta float64 = 0.001

var TestCases = []TestCase{
	{1, 1, 1, 0},
	{2, 1, 1, 0.5},
	{100000009, 33, 17, 1},
	{4, 2, 0, 0.25},
	{4, 2, 1, 0.5},
	{4, 2, 2, 0.25},
}

func TestCupTower(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := champagneTower(tc.Poured, tc.QRow, tc.QGlass)
		assert.InDelta(want, got, delta, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkCupTower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			champagneTower(tc.Poured, tc.QRow, tc.QGlass)
		}
	}
}
