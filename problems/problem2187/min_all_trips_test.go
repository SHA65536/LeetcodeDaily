package problem2187

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Trips    int
	Expected int64
}

var Results = []Result{
	{[]int{1, 2, 3}, 5, 3},
	{[]int{2}, 1, 2},
}

func TestMinimumTimeForAllTrips(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected

		got := minimumTime(res.Input, res.Trips)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
