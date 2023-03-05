package problem1345

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
	{[]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}, 3},
	{[]int{7}, 0},
	{[]int{7, 6, 9, 6, 9, 6, 9, 7}, 1},
	{[]int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13}, 3},
}

func TestJumpGameIV(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected

		got := minJumps(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
