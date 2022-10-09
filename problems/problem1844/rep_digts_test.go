package problem1844

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
	{"a1c1e1", "abcdef"},
	{"a1b2c3d4e", "abbdcfdhe"},
}

func TestReplaceDigits(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := replaceDigits(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
