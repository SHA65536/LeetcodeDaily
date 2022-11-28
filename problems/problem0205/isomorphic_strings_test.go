package problem0205

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Src, Dst string
	Expected bool
}

var Results = []Result{
	{"egg", "add", true},
	{"foo", "bar", false},
	{"paper", "title", true},
}

func TestIsomorphicStrings(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isIsomorphic(res.Src, res.Dst)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
