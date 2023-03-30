package problem0087

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	S1, S2   string
	Expected bool
}

var Results = []Result{
	{"great", "rgeat", true},
	{"abcde", "caebd", false},
	{"a", "a", true},
	{"acddaaaadbcbdcdaccabdbadccaaa", "adcbacccabbaaddadcdaabddccaaa", false},
}

func TestScrambleStringRec(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isScrambleRec(res.S1, res.S2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkScrambleStringRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			isScrambleRec(res.S1, res.S2)
		}
	}
}

func TestScrambleStringRecIter(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isScrambleIter(res.S1, res.S2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkScrambleStringIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			isScrambleIter(res.S1, res.S2)
		}
	}
}
