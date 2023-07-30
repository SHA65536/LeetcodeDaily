package problem0664

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    string
	Expected int
}

var TestCases = []TestCase{
	{"aaabbb", 2},
	{"aba", 2},
	{"abab", 3},
	{"aaa", 1},
	{"khigjjlkfjglgcegibdaiagkagfcajja", 21},
	{"jciacechlhlgilbelihfdhliihkakjbd", 21},
	{"hjjcgadfglbakfedhdkkakjkciblkifc", 22},
	{"effjefakiibfcgkidhgfjkkgllehefji", 22},
	{"jfccakhelfebacedbfjdafibigiecdeg", 23},
	{"hkdifcjaagjcjciejbcaailkbbjidifk", 20},
	{"ekgfhefkdebkbdfefidbgbbldcglaafa", 20},
	{"kbgcdfdaiafkiheilehhdlfdejeahhhd", 21},
}

func TestStrangePrinter(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := strangePrinter(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkStrangePrinter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			strangePrinter(tc.Input)
		}
	}
}
