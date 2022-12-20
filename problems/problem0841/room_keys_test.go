package problem0841

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected bool
}

var Results = []Result{
	{[][]int{{1}, {2}, {3}, {}}, true},
	{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
}

func TestCanVisitAllRooms(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := canVisitAllRooms(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
