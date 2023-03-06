package problem1539

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
	{[]int{2, 3, 4, 7, 11}, 5, 9},
	{[]int{1, 2, 3, 4}, 2, 6},
}

func TestKthMissingNumber(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected

		got := findKthPositive(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
