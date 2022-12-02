package problem1657

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Word1, Word2 string
	Expected     bool
}

var Results = []Result{
	{"abc", "bca", true},
	{"a", "aa", false},
	{"cabbba", "abbccc", true},
	{"cabbba", "cabbad", false},
	{"uau", "ssx", false},
}

func TestStringsAreSimilar(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := closeStrings(res.Word1, res.Word2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
