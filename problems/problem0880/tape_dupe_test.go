package problem0880

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    string
	K        int
	Expected string
}

var TestCases = []TestCase{
	{"leet2code3", 10, "o"},
	{"ha22", 5, "h"},
	{"a2345678999999999999999", 1, "a"},
}

func TestTapeDupe(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := decodeAtIndex(tc.Input, tc.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkTapeDupe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			decodeAtIndex(tc.Input, tc.K)
		}
	}
}
