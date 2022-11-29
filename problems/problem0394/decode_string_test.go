package problem0394

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"3[a]2[bc]", "aaabcbc"},
	{"3[a2[c]]", "accaccacc"},
	{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	{"3[abc2[3[av]]]", "abcavavavavavavabcavavavavavavabcavavavavavav"},
	{"3[3[3[3[a]]]]", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
}

func TestDecodeString(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := decodeString(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkDecodeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			decodeString(res.Input)
		}
	}
}

func TestDecodeStringRecursive(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := decodeStringRecursive(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkDecodeStringRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			decodeStringRecursive(res.Input)
		}
	}
}
