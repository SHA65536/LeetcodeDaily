package problem0461

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	A, B     int
	Expected int
}

var Results = []Result{
	{1, 4, 2},
	{3, 1, 1},
	{15, 78, 2},
	{15, 15, 0},
}

func TestHammingDistance(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		got := hammingDistance(res.A, res.B)
		want := res.Expected
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
