package problem2444

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input      []int
	MinK, MaxK int
	Expected   int64
}

var Results = []Result{
	{[]int{1, 3, 5, 2, 7, 5}, 1, 5, 2},
	{[]int{1, 1, 1, 1}, 1, 1, 10},
}

func TestBoundSubArrays(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected

		got := countSubarrays(res.Input, res.MinK, res.MaxK)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
