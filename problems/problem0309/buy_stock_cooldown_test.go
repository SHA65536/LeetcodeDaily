package problem0309

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
	{[]int{1, 2, 3, 0, 2}, 3},
	{[]int{1}, 0},
}

func TestBuyStockWithCooldown(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxProfit(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
