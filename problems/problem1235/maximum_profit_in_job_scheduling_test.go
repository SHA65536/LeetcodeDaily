package problem1235

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Start    []int
	End      []int
	Profit   []int
	Expected int
}

var Results = []Result{
	{[]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}, 120},
	{[]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}, 150},
	{[]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}, 6},
}

func TestMaxSchedulingProfit(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := jobScheduling(res.Start, res.End, res.Profit)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
