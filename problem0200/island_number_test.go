package problem0200

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]byte
	Expected int
}

var Results = []Result{
	{[][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}, 1},
	{[][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}, 3},
	{[][]byte{
		{'1'}}, 1},
	{[][]byte{
		{'0'}}, 0},
	{[][]byte{
		{'1', '0', '1', '0', '1'},
		{'0', '1', '0', '1', '0'},
		{'1', '0', '1', '0', '1'},
		{'0', '1', '0', '1', '0'}}, 10},
}

func TestNumberOfIslands(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numIslands(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
