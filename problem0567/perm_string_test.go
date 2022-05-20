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
