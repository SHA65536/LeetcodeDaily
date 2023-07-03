package problem0859

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	S, Goal  string
	Expected bool
}

var TestCases = []TestCase{
	{"ab", "ba", true},
	{"ab", "ab", false},
	{"aa", "aa", true},
}

func TestBuddyStrings(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := buddyStrings(tc.S, tc.Goal)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkBuddyStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			buddyStrings(tc.S, tc.Goal)
		}
	}
}
