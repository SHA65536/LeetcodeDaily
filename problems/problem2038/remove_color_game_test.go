package problem2038

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Answer   string
	Expected bool
}

var TestCases = []TestCase{
	{"AAABABB", true},
	{"AA", false},
	{"ABBBBBBBAAA", false},
}

func TestRemoveColorGame(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := winnerOfGame(tc.Answer)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkRemoveColorGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			winnerOfGame(tc.Answer)
		}
	}
}
