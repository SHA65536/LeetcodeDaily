package problem0933

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Cmds     []string
	Vals     []int
	Expected []int
}

var Results = []Result{
	{
		[]string{"RecentCounter", "ping", "ping", "ping", "ping"},
		[]int{0, 1, 100, 3001, 3002},
		[]int{0, 1, 2, 3, 3},
	},
	{
		func() []string {
			res := make([]string, 10001)
			for i := range res {
				res[i] = "ping"
			}
			res[0] = "RecentCounter"
			return res
		}(),
		func() []int {
			res := make([]int, 10001)
			for i := range res {
				res[i] = i
			}
			return res
		}(),
		func() []int {
			res := make([]int, 10001)
			for i := range res {
				res[i] = i
				if i >= 3001 {
					res[i] = 3001
				}
			}
			return res
		}(),
	},
}

func TestRecentPings(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]int, len(want))
		var Pinger RecentCounter
		for i := range res.Cmds {
			switch res.Cmds[i] {
			case "RecentCounter":
				Pinger = Constructor()
			case "ping":
				got[i] = Pinger.Ping(res.Vals[i])
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
