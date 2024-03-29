package problem0567

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	S1       string
	S2       string
	Expected bool
}

var Results = []Result{
	{"ab", "eidbaooo", true},
	{"ab", "eidboaoo", false},
	{"abcdefgasdwe", "eidboaoo", false},
	{"a", "eidboaoo", true},
	{"a", "ab", true},
	{"a", "ba", true},
	{"abdef", "e", false},
}

func TestCheckInclusion(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := checkInclusion(res.S1, res.S2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCheckInclusion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			checkInclusion(res.S1, res.S2)
		}
	}
}

func TestCheckInclusionOld(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := checkInclusionOld(res.S1, res.S2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCheckInclusionOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			checkInclusionOld(res.S1, res.S2)
		}
	}
}
