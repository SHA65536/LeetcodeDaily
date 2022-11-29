package problem0844

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	A, B     string
	Expected bool
}

var Results = []Result{
	{"ab#c", "ad#c", true},
	{"ab##", "c#d#", true},
	{"a#c", "b", false},
}

func TestBackSpaceString(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := backspaceCompare(res.A, res.B)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
