package problem0389

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	S, T     string
	Expected byte
}

var TestCases = []TestCase{
	{"abcd", "abcde", 'e'},
	{"abdc", "abcde", 'e'},
	{"", "y", 'y'},
}

func TestFindDifference(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := findTheDifference(tc.S, tc.T)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkFindDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			findTheDifference(tc.S, tc.T)
		}
	}
}
