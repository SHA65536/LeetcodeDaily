package problem1011

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Days     int
	Expected int
}

var Results = []Result{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
	{[]int{3, 2, 2, 4, 1, 4}, 3, 6},
	{[]int{1, 2, 3, 1, 1}, 4, 3},
}

func TestCapToShip(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := shipWithinDays(res.Input, res.Days)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
