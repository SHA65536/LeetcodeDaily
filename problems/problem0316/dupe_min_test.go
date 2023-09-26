package problem0316

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
	{"bcabc", "abc"},
	{"cbacdcbc", "acdb"},
	{"gaiidfahiihebeficegeibdeeacfaibi", "adfhbcegi"},
	{"bfhgeaedhgbfaidecceffgeebafaiihg", "abdcefgih"},
	{"hdifefbhfdfgbahdhffigedhbcbffhdb", "adfigebch"},
	{"ebggcecbdaddgbbfachbhfbaaebgcdhb", "abcfegdh"},
}

func TestRemoveDuplicateMin(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := removeDuplicateLetters(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkRemoveDuplicateMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			removeDuplicateLetters(tc.Input)
		}
	}
}

func TestRemoveDuplicateMinArr(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := removeDuplicateLettersArr(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkRemoveDuplicateMinArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			removeDuplicateLettersArr(tc.Input)
		}
	}
}
