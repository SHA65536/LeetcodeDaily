package problem1876

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
	{"xyzzaz", 1},
	{"aababcabc", 4},
	{"aababcabtklskeqxkbkfuhhhyuquqxlzxldpqcznylfbembczmwnhxtuuijuijkkyykhsvmjc", 56},
	{"jncpzlsxsavaatvbzmbolcnhletwtulxtuhkvuggrhipbxagvkeydlpwnrmptlia", 55},
	{"ftysgpgrhaxrkosojhbgfouglablqalnmhmpnccotmxyaaspoglwgbrgfuyydhzf", 53},
}

func TestDistinctTripleSubstring(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countGoodSubstrings(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
