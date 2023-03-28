package problem0983

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Days     []int
	Costs    []int
	Expected int
}

var Results = []Result{
	{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17},
	{[]int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 24, 25, 27, 28, 29, 30, 31, 34, 37, 38, 39, 41, 43, 44, 45, 47, 48, 49, 54, 57, 60, 62, 63, 66, 69, 70, 72, 74, 76, 78, 80, 81, 82, 83, 84, 85, 88, 89, 91, 93, 94, 97, 99}, []int{9, 38, 134}, 423},
}

func TestMinCostOfTickets(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := mincostTickets(res.Days, res.Costs)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
