package problem0990

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected bool
}

var Results = []Result{
	{[]string{"a==b", "b!=a"}, false},
	{[]string{"b==a", "a==b"}, true},
	{[]string{"b==a", "b==c", "c!=a"}, false},
	{[]string{"c==c", "b==d", "x!=z"}, true},
}

func TestSatisfiabilityOfEquations(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := equationsPossible(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
