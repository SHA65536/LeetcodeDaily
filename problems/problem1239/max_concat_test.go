package problem1239

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected int
}

var Results = []Result{
	{[]string{"un", "iq", "ue"}, 4},
	{[]string{"cha", "r", "act", "ers"}, 6},
	{[]string{"abcdefghijklmnopqrstuvwxyz"}, 26},
}

func TestMaxUniqueConcatLength(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxLength(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
