package problem0909

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
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1}}, 4},
	{[][]int{
		{-1, -1},
		{-1, 3}}, 1},
	{[][]int{
		{-1, 83, -1, 46, -1, -1, -1, -1, 40, -1},
		{-1, 29, -1, -1, -1, 51, -1, 18, -1, -1},
		{-1, 35, 31, 51, -1, 6, -1, 40, -1, -1},
		{-1, -1, -1, 28, -1, 36, -1, -1, -1, -1},
		{-1, -1, -1, -1, 44, -1, -1, 84, -1, -1},
		{-1, -1, -1, 31, -1, 98, 27, 94, 74, -1},
		{4, -1, -1, 46, 3, 14, 7, -1, 84, 67},
		{-1, -1, -1, -1, 2, 72, -1, -1, 86, -1},
		{-1, 32, -1, -1, -1, -1, -1, -1, -1, 19},
		{-1, -1, -1, -1, -1, 72, 46, -1, 92, 6}}, 3},
}

func TestSnakesAndLadders(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := snakesAndLadders(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
