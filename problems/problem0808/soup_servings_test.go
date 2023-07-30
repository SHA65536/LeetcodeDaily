package problem0808

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N        int
	Expected float64
}

const precision float64 = 0.00001

var TestCases = []TestCase{
	{50, .625},
	{100, .71875},
}

func TestSoupProbability(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := soupServings(tc.N)
		assert.InDelta(want, got, precision, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSoupProbability(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			soupServings(tc.N)
		}
	}
}
