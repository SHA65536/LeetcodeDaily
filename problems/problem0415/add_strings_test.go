package problem0415

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	A, B     string
	Expected string
}

var Results = []Result{
	{"2", "3", "5"},
	{"123", "456", "579"},
	{"123456", "654321", "777777"},
}

func TestAddStrings(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := addStrings(res.A, res.B)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
