package problem0118

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected [][]int
}

var Results = []Result{
	{1, [][]int{{1}}},
	{2, [][]int{{1}, {1, 1}}},
	{3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
	{4, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
	{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
}

func TestGeneratePascal(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := generate(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
