package problem1383

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N                 int
	Speed, Efficiency []int
	K                 int
	Expected          int
}

var Results = []Result{
	{
		6,
		[]int{2, 10, 3, 1, 5, 8},
		[]int{5, 4, 3, 9, 7, 2},
		2,
		60,
	},
	{
		6,
		[]int{2, 10, 3, 1, 5, 8},
		[]int{5, 4, 3, 9, 7, 2},
		3,
		68,
	},
	{
		6,
		[]int{2, 10, 3, 1, 5, 8},
		[]int{5, 4, 3, 9, 7, 2},
		4,
		72,
	},
}

func TestMaxTeamPerformance(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxPerformance(res.N, res.Speed, res.Efficiency, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
