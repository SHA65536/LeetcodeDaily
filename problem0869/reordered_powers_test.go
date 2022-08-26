package problem0869

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
	{10, false},
	{8, true},
	{15, false},
	{16, true},
	{61, true},
	{66, false},
	{4082, true},
}

func TestReorderedPowerOf2(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reorderedPowerOf2(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
