package problem1465

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	H, W     int
	Hc, Vc   []int
	Expected int
}

var Results = []Result{
	{5, 4, []int{1, 2, 4}, []int{1, 3}, 4},
	{5, 4, []int{3, 1}, []int{1}, 6},
	{5, 4, []int{3}, []int{3}, 9},
}

func TestMaxCakeArea(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxArea(res.H, res.W, res.Hc, res.Vc)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
