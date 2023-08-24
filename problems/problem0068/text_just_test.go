package problem0068

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    []string
	Width    int
	Expected []string
}

var TestCases = []TestCase{
	{
		[]string{"This", "is", "an", "example", "of", "text", "justification."},
		16,
		[]string{"This    is    an",
			"example  of text",
			"justification.  "},
	},
	{
		[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
		16,
		[]string{"What   must   be",
			"acknowledgment  ",
			"shall be        "},
	},
	{
		[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
		20,
		[]string{"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  "}},
}

func TestJustifyText(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := fullJustify(tc.Input, tc.Width)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkJustifyText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			fullJustify(tc.Input, tc.Width)
		}
	}
}
