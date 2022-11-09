package problem0058

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected int
}

var Results = []Result{
	{"Hello World", 5},
	{"   fly me   to   the moon  ", 4},
	{"luffy is still joyboy", 6},
	{"a", 1},
	{" a", 1},
	{"a ", 1},
}

func TestSLengthOfLastWord(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := lengthOfLastWord(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkLenthOfLastWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			lengthOfLastWord(res.Input)
		}
	}
}

func TestSLengthOfLastWordSplit(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := lengthOfLastWordSplit(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkLenthOfLastWordSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			lengthOfLastWordSplit(res.Input)
		}
	}
}
