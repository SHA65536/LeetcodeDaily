package problem1845

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Cmds     []string
	Vals     []int
	Expected []int
}

var TestCases = []TestCase{
	{
		[]string{"SeatManager", "reserve", "reserve", "unreserve", "reserve", "reserve", "reserve", "reserve", "unreserve"},
		[]int{5, 0, 0, 2, 0, 0, 0, 0, 5},
		[]int{0, 1, 2, 0, 2, 3, 4, 5, 0},
	},
}

func TestSeatManager(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var want = tc.Expected
		var got = make([]int, len(want))
		var manager SeatManager
		for i, cmd := range tc.Cmds {
			switch cmd {
			case "SeatManager":
				manager = Constructor(tc.Vals[i])
			case "reserve":
				got[i] = manager.Reserve()
			case "unreserve":
				manager.Unreserve(tc.Vals[i])
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSeatManager(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			var manager SeatManager
			for i, cmd := range tc.Cmds {
				switch cmd {
				case "SeatManager":
					manager = Constructor(tc.Vals[i])
				case "reserve":
					manager.Reserve()
				case "unreserve":
					manager.Unreserve(tc.Vals[i])
				}
			}
		}
	}
}
