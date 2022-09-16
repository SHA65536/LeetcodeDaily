package problem0063

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected int
}

var Results = []Result{
	{[][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0}}, 2},
	{[][]int{
		{0, 1},
		{0, 0}}, 1},
	{[][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0}}, 6},
	{[][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}, 20},
}

func TestUniquePathsWithObstacles(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := uniquePathsWithObstacles(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
