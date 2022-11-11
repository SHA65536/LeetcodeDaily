package problem0069

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
	{1, 1},
	{2, 1},
	{3, 1},
	{4, 2},
	{8, 2},
	{9, 3},
	{10, 3},
}

func TestRemoveAdjacent(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := mySqrt(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
