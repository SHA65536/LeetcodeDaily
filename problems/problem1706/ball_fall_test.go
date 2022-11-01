package problem1706

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected []int
}

var Results = []Result{
	{[][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}, []int{1, -1, -1, -1, -1}},
	{[][]int{{-1}}, []int{-1}},
	{[][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}, []int{0, 1, 2, 3, 4, -1}},
}

func TestBallsFalling(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findBall(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
