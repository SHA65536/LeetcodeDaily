package problem0242

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Target   string
	Expected bool
}

var Results = []Result{
	{"anagram", "nagaram", true},
	{"cat", "car", false},
	{"abc", "ab", false},
	{"abc", "abc", true},
}

func TestIsAnagram(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isAnagram(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
