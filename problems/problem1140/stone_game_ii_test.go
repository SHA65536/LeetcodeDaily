package problem1140

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected int
}

var Results = []Result{
	{[]int{2, 7, 9, 4, 4}, 10},
	{[]int{1, 2, 3, 4, 5, 100}, 104},
}

func TestStoneGameIIDP(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := stoneGameIIDP(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkStoneGameIIDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			stoneGameIIDP(res.Input)
		}
	}
}

func TestStoneGameIIMiniMax(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := stoneGameIIMiniMax(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkStoneGameIIMiniMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			stoneGameIIMiniMax(res.Input)
		}
	}
}
