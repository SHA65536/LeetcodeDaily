package problem2413

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
	{5, 10},
	{6, 6},
	{31, 62},
	{32, 32},
}

func TestSmallestEvenMultiple(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := smallestEvenMultiple(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
