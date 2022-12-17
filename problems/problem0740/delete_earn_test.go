package problem0740

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
	{[]int{3, 4, 2}, 6},
	{[]int{2, 2, 3, 3, 3, 4}, 9},
}

func TestDeleteAndEarn(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := deleteAndEarn(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
