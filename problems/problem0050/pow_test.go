package problem0050

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	X        float64
	N        int
	Expected float64
}

const precision float64 = 0.00001

var TestCases = []TestCase{
	{2, 10, math.Pow(2, 10)},
	{2.1, 3, math.Pow(2.1, 3)},
	{2, -2, math.Pow(2, -2)},
	{57, 1547, math.Pow(57, 1547)},
	{57, -1547, math.Pow(57, -1547)},
}

func TestPow(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := myPow(tc.X, tc.N)
		assert.InDelta(want, got, precision, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			myPow(tc.X, tc.N)
		}
	}
}

func TestMathPow(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := math.Pow(tc.X, float64(tc.N))
		assert.InDelta(want, got, precision, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMathPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			math.Pow(tc.X, float64(tc.N))
		}
	}
}
