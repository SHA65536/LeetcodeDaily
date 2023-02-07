package problem0904

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
	{[]int{1, 2, 1}, 3},
	{[]int{0, 1, 2, 2}, 3},
	{[]int{1, 2, 3, 2, 2}, 4},
}

func TestTotalFruitPicked(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := totalFruit(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
