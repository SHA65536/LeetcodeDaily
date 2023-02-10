package problem1496

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
	{"NES", false},
	{"NESWW", true},
}

func TestPathCrossingItself(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isPathCrossing(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
