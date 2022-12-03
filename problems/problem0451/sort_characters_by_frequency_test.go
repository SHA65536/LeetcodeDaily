package problem0451

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
	{"tree", "eetr"},
	{"cccaaa", "cccaaa"},
	{"Aabb", "bbaA"},
	{"loveleetcode", "eeeeoollvtdc"},
	{"2a554442f544asfasssffffasss", "sssssssffffff44444aaaa55522"},
}

func TestSortFreq(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := frequencySort(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSortFreq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			frequencySort(res.Input)
		}
	}
}

func TestSortFreqNoMap(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := frequencySortNoMap(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSortFreqNoMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			frequencySortNoMap(res.Input)
		}
	}
}
