package problem1391

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Roads    [][]int
	Expected bool
}

var Results = []Result{
	{[][]int{{2, 4, 3}, {6, 5, 2}}, true},
	{[][]int{{1, 2, 1}, {1, 2, 1}}, false},
	{[][]int{{1, 1, 2}}, false},
	{[][]int{
		{4, 1, 3},
		{6, 1, 2},
	}, true},
	{[][]int{
		{4, 1, 3},
		{6, 1, 3},
	}, true},
	{[][]int{{1, 1, 1, 1, 1, 1, 3}}, true},
	{[][]int{{1, 2}, {2, 1}}, false},
	{[][]int{{4, 1}, {6, 1}}, true},
}

func TestValidRoadpath(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := hasValidPath(res.Roads)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkValidRoadPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			hasValidPath(res.Roads)
		}
	}
}
