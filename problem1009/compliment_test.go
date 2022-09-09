package problem1009

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected int
}

var Results = []Result{
	{5, 2},
	{1, 0},
	{7, 0},
	{8, 7},
	{0, 1},
}

func TestNumberCompliment(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := bitwiseComplement(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
