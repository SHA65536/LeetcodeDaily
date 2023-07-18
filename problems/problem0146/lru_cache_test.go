package problem0146

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Commands []string
	Values   [][]int
	Expected []int
}

var TestCases = []TestCase{
	{
		[]string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
		[][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
		[]int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4},
	},
}

func TestAddTwoNumbersII(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := make([]int, len(want))
		var cache LRUCache
		for i := range tc.Commands {
			switch tc.Commands[i] {
			case "LRUCache":
				cache = Constructor(tc.Values[i][0])
			case "put":
				cache.Put(tc.Values[i][0], tc.Values[i][1])
			case "get":
				got[i] = cache.Get(tc.Values[i][0])
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkAddTwoNumbersII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			want := tc.Expected
			got := make([]int, len(want))
			for i := range tc.Commands {
				var cache LRUCache
				switch tc.Commands[i] {
				case "LRUCache":
					cache = Constructor(tc.Values[i][0])
				case "put":
					cache.Put(tc.Values[i][0], tc.Values[i][1])
				case "get":
					got[i] = cache.Get(tc.Values[i][0])
				}
			}
		}
	}
}
