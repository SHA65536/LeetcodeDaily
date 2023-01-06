package problem1833

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Coins    int
	Expected int
}

var Results = []Result{
	{[]int{1, 3, 2, 4, 1}, 7, 4},
	{[]int{10, 6, 8, 7, 7, 8}, 5, 0},
	{[]int{1, 6, 3, 1, 2, 5}, 20, 6},
}

func TestMaxIceCreamBars(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxIceCream(res.Input, res.Coins)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
