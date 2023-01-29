package problem0460

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Vals     [][]int
	Expected []int
}

var Results = []Result{
	{
		[]string{"LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"},
		[][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {3}, {4, 4}, {1}, {3}, {4}},
		[]int{0, 0, 0, 1, 0, -1, 3, 0, -1, 3, 4},
	},
}

func TestLFUCache(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]int, len(res.Expected))
		var lfu LFUCache
		for i := range res.Input {
			switch res.Input[i] {
			case "LFUCache":
				lfu = Constructor(res.Vals[i][0])
			case "put":
				lfu.Put(res.Vals[i][0], res.Vals[i][1])
			case "get":
				got[i] = lfu.Get(res.Vals[i][0])
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
