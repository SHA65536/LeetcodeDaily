package problem0997

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N        int
	Trust    [][]int
	Expected int
}

var Results = []Result{
	{2, [][]int{{1, 2}}, 2},
	{3, [][]int{{1, 3}, {2, 3}}, 3},
	{3, [][]int{{1, 3}, {2, 3}, {3, 1}}, -1},
}

func TestRestoreAddresses(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findJudge(res.N, res.Trust)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
