package problem1357

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N, Discount      int
	Products, Prices []int
	Buy, Amount      [][]int
	Expected         []float64
}

var Results = []Result{
	{
		N: 3, Discount: 50,
		Products: []int{1, 2, 3, 4, 5, 6, 7},
		Prices:   []int{100, 200, 300, 400, 300, 200, 100},
		Buy:      [][]int{{1, 2}, {3, 7}, {1, 2, 3, 4, 5, 6, 7}, {4}, {7, 3}, {7, 5, 3, 1, 6, 4, 2}, {2, 3, 5}},
		Amount:   [][]int{{1, 2}, {10, 10}, {1, 1, 1, 1, 1, 1, 1}, {10}, {10, 10}, {10, 10, 10, 9, 9, 9, 7}, {5, 3, 2}},
		Expected: []float64{500.0, 4000.0, 800.0, 4000.0, 4000.0, 7350.0, 2500.0},
	},
}

func TestDiscoutnCashier(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]float64, len(want))
		cashier := Constructor(res.N, res.Discount, res.Products, res.Prices)
		for i := range res.Expected {
			got[i] = cashier.GetBill(res.Buy[i], res.Amount[i])
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkDiscoutnCashier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			cashier := Constructor(res.N, res.Discount, res.Products, res.Prices)
			for i := range res.Expected {
				cashier.GetBill(res.Buy[i], res.Amount[i])
			}
		}
	}
}
