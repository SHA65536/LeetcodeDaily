package problem1254

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
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0}}, 2},
	{[][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0}}, 1},
	{[][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1}}, 2},
	{[][]int{
		{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 0}}, 5},
}

func TestClosedIslands(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := closedIsland(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
