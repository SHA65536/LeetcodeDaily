package problem0020

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
	{"([]{})", true},
	{"()()", true},
	{"([{}])", true},
	{"([}{])", false},
	{")[]{})", false},
	{"((())", false},
	{"()", true},
}

func TestIsValidParenthesis(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		assert.Equal(want, isValid(res.Input), fmt.Sprintf("%+v", res))
	}
}
