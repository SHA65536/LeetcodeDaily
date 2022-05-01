package problem0008

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected int
}

var Results = []Result{
	{"42", 42},
	{"   -42", -42},
	{"4193 with words", 4193},
	{"2147483647", 2147483647},
	{"21407483647", 2147483647},
	{"02147483647", 2147483647},
	{"  99999999999abc", 2147483647},
	{"-2147483647", -2147483647},
	{"-2147483648", -2147483648},
	{"-2147483649", -2147483648},
	{"-00000002147483649", -2147483648},
}

func TestMyAtoi(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := myAtoi(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
