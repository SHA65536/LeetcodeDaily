package problem1759

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
	{"abbcccaa", 13},
	{"xy", 2},
	{"zzzzz", 15},
}

func TestCountHomogenous(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var want = tc.Expected
		got := countHomogenous(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkCountHomogenous(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			countHomogenous(tc.Input)
		}
	}
}
