package problem0886

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N        int
	Input    [][]int
	Expected bool
}

var Results = []Result{
	{4, [][]int{{1, 2}, {1, 3}, {2, 4}}, true},
	{3, [][]int{{1, 2}, {1, 3}, {2, 3}}, false},
	{5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}, false},
}

func TestPossibleBipartition(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := possibleBipartition(res.N, res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPossibleBipartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			possibleBipartition(res.N, res.Input)
		}
	}
}
