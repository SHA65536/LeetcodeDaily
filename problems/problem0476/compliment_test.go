package problem0476

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
}

func TestNumberCompliment(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findComplement(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
