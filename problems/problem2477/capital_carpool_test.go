package problem2477

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Seats    int
	Expected int64
}

var Results = []Result{
	{
		[][]int{{0, 1}, {0, 2}, {0, 3}}, 5,
		3,
	},
	{
		[][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2,
		7,
	},
	{
		[][]int{}, 1,
		0,
	},
}

func TestMinFuelCarpool(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minimumFuelCost(res.Input, res.Seats)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
