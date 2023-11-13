package problem2785

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    string
	Expected string
}

var TestCases = []TestCase{
	{"lEetcOde", "lEOtcede"},
	{"lYmpH", "lYmpH"},
	{"cAtmRFPcRKnizPOAVjvZkDXgulkviIPi", "cAtmRFPcRKnAzPIOVjvZkDXgilkviiPu"},
	{"SXbKQzUkqkWlaNcHBGcovKYTTRoCCwRZ", "SXbKQzUkqkWlaNcHBGcovKYTTRoCCwRZ"},
}

func TestSortVowelsInAString(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var want = tc.Expected
		var got = sortVowels(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSortVowelsInAString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			sortVowels(tc.Input)
		}
	}
}

func TestSortVowelsInAStringOpt(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var want = tc.Expected
		var got = sortVowelsOpt(tc.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSortVowelsInAStringOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			sortVowelsOpt(tc.Input)
		}
	}
}
