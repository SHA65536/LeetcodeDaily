package problem0779

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N        int
	K        int
	Expected int
}

var TestCases = []TestCase{
	{1, 1, 0},
	{2, 1, 0},
	{2, 2, 1},
	{30, 434991989, 0},
}

func TestKthSymbol(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := kthGrammar(tc.N, tc.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkKthSymbol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			kthGrammar(tc.N, tc.K)
		}
	}
}
