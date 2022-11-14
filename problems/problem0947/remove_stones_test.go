package problem0947

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
		[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}},
		5,
	},
	{
		[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}},
		3,
	},
	{
		[][]int{{0, 0}},
		0,
	},
}

func TestRemoveStonesDFS(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := removeStonesDFS(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkRemoveStonesDFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			removeStonesDFS(res.Input)
		}
	}
}

func TestRemoveStonesUF(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := removeStonesUF(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkRemoveStonesUF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			removeStonesUF(res.Input)
		}
	}
}

/*

 */
