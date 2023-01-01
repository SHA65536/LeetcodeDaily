package problem0290

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Pat, S   string
	Expected bool
}

var Results = []Result{
	{"abba", "dog cat cat dog", true},
	{"abba", "dog cat cat fish", false},
	{"aaaa", "dog cat cat dog", false},
	{"abcd", "dog cat cat dog", false},
}

func TestWordPattern(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := wordPattern(res.Pat, res.S)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
