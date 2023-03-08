package problem0875

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Time     int
	Expected int
}

var Results = []Result{
	{[]int{3, 6, 7, 11}, 8, 4},
	{[]int{30, 11, 23, 4, 20}, 5, 30},
	{[]int{30, 11, 23, 4, 20}, 6, 23},
}

func TestKokoBananas(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected

		got := minEatingSpeed(res.Input, res.Time)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
