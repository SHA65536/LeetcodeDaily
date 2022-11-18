package problem0263

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected bool
}

var Results = []Result{
	{1, true},
	{6, true},
	{14, false},
	{35, false},
	{36, true},
}

func TestIsUglyNumber(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isUgly(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
