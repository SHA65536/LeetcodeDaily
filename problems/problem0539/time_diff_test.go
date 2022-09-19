package problem0539

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected int
}

var Results = []Result{
	{[]string{"23:59", "00:00"}, 1},
	{[]string{"00:00", "23:59", "00:00"}, 0},
	{[]string{"23:59", "01:00"}, 61},
	{[]string{"01:00", "02:34"}, 94},
}

func TestMinimumTimeDifference(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findMinDifference(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
