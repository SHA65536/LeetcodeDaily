package problem0871

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Target, Start int
	Stations      [][]int
	Expected      int
}

var Results = []Result{
	{1, 1, [][]int{}, 0},
	{100, 1, [][]int{{10, 100}}, -1},
	{100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}, 2},
	{100, 50, [][]int{{25, 50}, {50, 25}}, 1},
}

func TestRefuelStops(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minRefuelStops(res.Target, res.Start, res.Stations)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
