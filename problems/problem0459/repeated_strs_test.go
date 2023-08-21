package problem0459

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    string
	Expected bool
}

var TestCases = []TestCase{
	{"abab", true},
	{"aba", false},
	{"abcabcabcabc", true},
	{"abcdabcdabcdabcdabcdabcdabcdabcdabcd", true},
	{"abcdabcdabcdabcdabcdabcdabcdabcdabcde", false},
}

func TestRepeatedStrings(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := repeatedSubstringPattern(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkRepeatedStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			repeatedSubstringPattern(tc.Input)
		}
	}
}
