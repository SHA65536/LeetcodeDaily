package problem0785

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
	{[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}, false},
	{[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}, true},
}

func TestIsGraphBipartite(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isBipartite(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkIsGraphBipartite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			isBipartite(res.Input)
		}
	}
}
