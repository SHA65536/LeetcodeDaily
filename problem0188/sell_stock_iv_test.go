package problem0188

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	K        int
	Prices   []int
	Expected int
}

var Results = []Result{
	{2, []int{2, 4, 1}, 2},
	{2, []int{3, 2, 6, 5, 0, 3}, 7},
	{10, []int{91, 7, 5, 28, 69, 70, 3, 5, 92, 26, 13, 7, 51, 39, 74, 1, 58, 14, 87, 44, 15, 53, 57, 71, 3, 76, 66, 89, 33, 20, 23, 70, 81, 15, 9, 84, 72, 0, 50, 51, 7, 4, 96, 53, 3, 13, 7, 86, 81, 47, 9, 4, 71, 78, 94, 2, 30, 51, 56, 59, 66, 71, 1, 0, 36, 10, 45, 15, 26, 84, 44, 53, 70, 29, 62, 51, 60, 78, 53, 88, 88, 39, 83, 18, 0, 43, 76, 65, 12, 41, 51, 44, 20, 32, 84, 17, 40, 78, 22, 25}, 846},
}

func TestSellStockIV(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxProfit(res.K, res.Prices)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
