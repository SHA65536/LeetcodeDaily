package problem0150

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
	{[]string{"2", "1", "+", "3", "*"}, 9},
	{[]string{"4", "13", "5", "/", "+"}, 6},
	{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
}

func TestEvalPostfix(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := evalRPN(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
