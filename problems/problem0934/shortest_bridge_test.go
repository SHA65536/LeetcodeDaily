package problem0934

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
	{
		[][]int{
			{0, 1},
			{1, 0}},
		1,
	},
	{
		[][]int{
			{0, 1, 0},
			{0, 0, 0},
			{0, 0, 1}},
		2,
	},
	{
		[][]int{
			{1, 1, 1, 1, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 1, 1, 1, 1}},
		1,
	},
}

func TestShortestBridge(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := shortestBridge(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkShortestBridge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			shortestBridge(res.Input)
		}
	}
}
