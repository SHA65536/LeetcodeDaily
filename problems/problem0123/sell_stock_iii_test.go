package problem0123

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
	{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
	{[]int{1, 2, 3, 4, 5}, 4},
	{[]int{7, 6, 4, 3, 1}, 0},
}

func TestSellStockIII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxProfit(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
