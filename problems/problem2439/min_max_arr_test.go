package problem2439

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
	{[]int{3, 7, 1, 6}, 5},
	{[]int{10, 1}, 10},
}

func TestMinMaxArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minimizeArrayValue(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
