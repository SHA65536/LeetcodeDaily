package problem0744

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []byte
	Target   byte
	Expected byte
}

var TestCases = []TestCase{
	{[]byte("cfj"), 'a', 'c'},
	{[]byte("xxyy"), 'z', 'x'},
	{[]byte("cfj"), 'c', 'f'},
	{[]byte("abcdefghijklmnopqrstuvwxyz"), 'k', 'l'},
	{[]byte("abcdefghijlmnopqrstuvwxyz"), 'k', 'l'},
	{[]byte("abcdefghijkmnopqrstuvwxyz"), 'k', 'm'},
}

func TestFindLetter(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := nextGreatestLetter(tc.Input, tc.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkFindLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			nextGreatestLetter(tc.Input, tc.Target)
		}
	}
}
