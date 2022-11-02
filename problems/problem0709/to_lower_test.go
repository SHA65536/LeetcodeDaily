package problem0709

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
	{"hello", "hello"},
	{"hERE", "here"},
	{"boop Beep", "boop beep"},
}

func TestToLowercase(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := toLowerCase(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
