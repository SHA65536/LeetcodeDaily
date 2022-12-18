package problem0739

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected []int
}

var Results = []Result{
	{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
	{[]int{30, 60, 90}, []int{1, 1, 0}},
}

func TestDailyTemperature(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := dailyTemperatures(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
