package problem1047

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
	{"abbaca", "ca"},
	{"azxxzy", "ay"},
}

func TestRemoveAdjacent(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := removeDuplicates(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
