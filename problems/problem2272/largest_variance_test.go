package problem2272

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    string
	Expected int
}

var TestCases = []TestCase{
	{"aababbb", 3},
	{"abcde", 0},
}

func TestLargestVariance(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := largestVariance(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkLargestVariance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			largestVariance(tc.Input)
		}
	}
}
