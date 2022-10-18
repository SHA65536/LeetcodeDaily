package problem0038

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected string
}

var Results = []Result{
	{1, "1"},
	{2, "11"},
	{3, "21"},
	{4, "1211"},
	{5, "111221"},
	{6, "312211"},
	{7, "13112221"},
	{8, "1113213211"},
	{10, "13211311123113112211"},
}

func TestCountAndSay(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countAndSay(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
