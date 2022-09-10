package problem0122

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
	{[]int{7, 1, 5, 3, 6, 4}, 7},
	{[]int{1, 2, 3, 4, 5}, 4},
	{[]int{7, 6, 4, 3, 1}, 0},
}

func TestSellStockII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxProfit(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
