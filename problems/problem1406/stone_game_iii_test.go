package problem1406

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected string
}

var Results = []Result{
	{[]int{1, 2, 3, 7}, "Bob"},
	{[]int{1, 2, 3, -9}, "Alice"},
	{[]int{1, 2, 3, 6}, "Tie"},
}

func TestStoneGameIII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := stoneGameIII(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkStoneGameIII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			stoneGameIII(res.Input)
		}
	}
}
