package problem1996

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected int
}

var Results = []Result{
	{[][]int{{5, 5}, {6, 3}, {3, 6}}, 0},
	{[][]int{{2, 2}, {3, 3}}, 1},
	{[][]int{{1, 5}, {10, 4}, {4, 3}}, 1},
}

func TestWeakestCharacters(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numberOfWeakCharacters(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
