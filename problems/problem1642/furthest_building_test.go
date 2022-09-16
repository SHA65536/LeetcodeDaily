package problem1642

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Heights         []int
	Bricks, Ladders int
	Expected        int
}

var Results = []Result{
	{
		[]int{4, 2, 7, 6, 9, 14, 12},
		5, 1, 4,
	},
	{
		[]int{14, 3, 19, 3},
		17, 0, 3,
	},
	{
		[]int{4, 12, 2, 7, 3, 18, 20, 3, 19},
		10, 2, 7,
	},
}

func TestFurthestBuilding(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := furthestBuilding(res.Heights, res.Bricks, res.Ladders)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
