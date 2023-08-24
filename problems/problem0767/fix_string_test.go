package problem0767

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    string
	Expected string
}

var TestCases = []TestCase{
	{"aab", "aba"},
	{"aaab", ""},
}

func TestReorganizeString(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := reorganizeString(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkReorganizeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			reorganizeString(tc.Input)
		}
	}
}
