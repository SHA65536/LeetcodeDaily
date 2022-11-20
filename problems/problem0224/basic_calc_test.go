package problem0224

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
	{"1 + 1", 2},
	{" 2-1 + 2 ", 3},
	{"(1+(4+5+2)-3)+(6+8)", 23},
	{"-1", -1},
}

func TestBasicCalculator(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := calculate(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
