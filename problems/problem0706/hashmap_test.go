package problem0706

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Cmds     []string
	Vals     [][]int
	Expected []int
}

var TestCases = []TestCase{
	{
		Cmds:     []string{"put", "put", "get", "get", "put", "get", "remove", "get"},
		Vals:     [][]int{{1, 1}, {2, 2}, {1}, {3}, {2, 1}, {2}, {2}, {2}},
		Expected: []int{0, 0, 1, -1, 0, 1, 0, -1},
	},
}

func TestHashMap(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := make([]int, len(tc.Expected))
		m := Constructor()
		for i, cmd := range tc.Cmds {
			switch cmd {
			case "put":
				m.Put(tc.Vals[i][0], tc.Vals[i][1])
			case "get":
				got[i] = m.Get(tc.Vals[i][0])
			case "remove":
				m.Remove(tc.Vals[i][0])
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkHashMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			m := Constructor()
			for i, cmd := range tc.Cmds {
				switch cmd {
				case "put":
					m.Put(tc.Vals[i][0], tc.Vals[i][1])
				case "get":
					m.Get(tc.Vals[i][0])
				case "remove":
					m.Remove(tc.Vals[i][0])
				}
			}
		}
	}
}
