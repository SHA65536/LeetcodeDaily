package problem1523

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Low, High int
	Expected  int
}

var Results = []Result{
	{3, 7, 3},
	{8, 10, 1},
	{1, 10, 5},
	{1, 1, 1},
	{2, 2, 0},
	{2, 4, 1},
	{1, 4, 2},
}

func TestMinFuelCarpool(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countOdds(res.Low, res.High)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
