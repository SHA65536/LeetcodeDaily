package problem0605

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	K        int
	Expected bool
}

var Results = []Result{
	{[]int{1, 0, 0, 0, 1}, 1, true},
	{[]int{1, 0, 0, 0, 1}, 2, false},
}

func TestCanPlaceFlowers(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := canPlaceFlowers(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
