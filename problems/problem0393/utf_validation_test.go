package problem0393

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected bool
}

var Results = []Result{
	{[]int{197, 130, 1}, true},
	{[]int{197, 130, 1, 197, 130, 1, 197, 130, 1}, true},
	{[]int{235, 140, 4}, false},
}

func TestValidUTF8(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := validUtf8(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
