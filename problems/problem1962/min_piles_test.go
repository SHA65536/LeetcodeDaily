package problem1962

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
	{[]int{5, 4, 9}, 2, 12},
	{[]int{4, 3, 6, 7}, 3, 12},
}

func TestMinStoneTotal(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minStoneSum(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
