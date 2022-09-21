package problem0985

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Queries  [][]int
	Expected []int
}

var Results = []Result{
	{[]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}, []int{8, 6, 2, 4}},
	{[]int{1}, [][]int{{4, 0}}, []int{0}},
}

func TestEvenSumAfterQueries(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := sumEvenAfterQueries(res.Input, res.Queries)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
