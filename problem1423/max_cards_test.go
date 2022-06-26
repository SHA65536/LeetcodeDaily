package problem1423

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	K        int
	Expected int
}

var Results = []Result{
	{[]int{1, 2, 3, 4, 5, 6, 1}, 3, 12},
	{[]int{2, 2, 2}, 2, 4},
	{[]int{9, 7, 7, 9, 7, 7, 9}, 7, 55},
}

func TestMaxScore(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxScore(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
