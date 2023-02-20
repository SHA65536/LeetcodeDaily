package problem0202

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
	{19, true},
	{2, false},
}

func TestHappyNumber(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isHappy(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
