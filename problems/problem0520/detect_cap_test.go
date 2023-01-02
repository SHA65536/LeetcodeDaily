package problem0520

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected bool
}

var Results = []Result{
	{"USA", true},
	{"leetcode", true},
	{"USa", false},
	{"Usa", true},
	{"UsA", false},
}

func TestDetectCapial(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := detectCapitalUse(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
