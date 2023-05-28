package problem1547

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N        int
	Cuts     []int
	Expected int
}

var Results = []Result{
	{7, []int{1, 3, 4, 5}, 16},
	{9, []int{5, 6, 1, 4, 2}, 22},
	{100, []int{1, 5, 7, 25, 75, 60, 24, 78, 63, 11, 15, 42}, 338},
}

func TestMinCostCutStick(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		temp := make([]int, len(res.Cuts))
		copy(temp, res.Cuts)
		got := minCost(res.N, temp)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMinCostCutStick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			temp := make([]int, len(res.Cuts))
			copy(temp, res.Cuts)
			minCost(res.N, temp)
		}
	}
}
