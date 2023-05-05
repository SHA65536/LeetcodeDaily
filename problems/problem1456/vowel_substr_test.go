package problem1456

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	K        int
	Expected int
}

var Results = []Result{
	{"abciiidef", 3, 3},
	{"aeiou", 2, 2},
	{"leetcode", 3, 2},
	{"aaa", 3, 3},
	{"bbb", 3, 0},
	{"a", 1, 1},
}

func TestMaxVowelSubstr(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxVowels(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxVowelSubstr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maxVowels(res.Input, res.K)
		}
	}
}
