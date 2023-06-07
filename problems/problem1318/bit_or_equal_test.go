package problem1318

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	A, B, C  int
	Expected int
}

var TestCases = []TestCase{
	{2, 6, 5, 3},
	{4, 2, 7, 1},
	{1, 2, 3, 0},
}

func TestFlipsToMakeOrEqual(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minFlips(tc.A, tc.B, tc.C)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkFlipsToMakeOrEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minFlips(tc.A, tc.B, tc.C)
		}
	}
}
