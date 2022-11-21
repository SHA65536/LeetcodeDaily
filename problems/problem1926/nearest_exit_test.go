package problem1926

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]byte
	Entrance []int
	Expected int
}

var Results = []Result{
	{
		[][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}},
		[]int{1, 2},
		1,
	},
	{
		[][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}},
		[]int{1, 0},
		2,
	},
	{
		[][]byte{{'.', '+'}},
		[]int{0, 0},
		-1,
	},
	{
		[][]byte{
			{'.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.'},
		},
		[]int{3, 3},
		3,
	},
	{
		[][]byte{
			{'+', '.', '+', '+', '+', '+', '+'},
			{'+', '.', '+', '.', '.', '.', '+'},
			{'+', '.', '+', '.', '+', '.', '+'},
			{'+', '.', '.', '.', '+', '.', '+'},
			{'+', '+', '+', '+', '+', '.', '+'},
		},
		[]int{0, 1},
		12,
	},
}

func TestNearestMazeExit(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := nearestExit(res.Input, res.Entrance)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
