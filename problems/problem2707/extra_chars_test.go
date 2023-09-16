package problem2707

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Str      string
	Dict     []string
	Expected int
}

var TestCases = []TestCase{
	{"leetscode", []string{"leet", "code", "leetcode"}, 1},
	{"sayhelloworld", []string{"hello", "world"}, 3},
}

func TestExtraStringChars(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minExtraChar(tc.Str, tc.Dict)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkExtraStringChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minExtraChar(tc.Str, tc.Dict)
		}
	}
}
