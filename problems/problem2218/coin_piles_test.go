package problem2218

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Piles    [][]int
	K        int
	Expected int
}

var Results = []Result{
	{[][]int{{1, 100, 3}, {7, 8, 9}}, 2, 101},
	{[][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7, 706},
	{[][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 550}}, 7, 601},
}

func TestCoinPiles(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxValueOfCoins(res.Piles, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
