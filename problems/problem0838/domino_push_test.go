package problem0838

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"RR.L", "RR.L"},
	{"RR.L", "RR.L"},
	{".L.R...LR..L..", "LL.RR.LLRRLL.."},
	{".LRL.LRL.LRRRL..R..LRR.LRR.RLLRLLLL..L.L.RR.LLR.LRLRLLRR.RL.LR..RR.L.R..R...RLRL.R.L.....RR.LL.LL.LL", "LLRLLLRLLLRRRL..RRLLRR.LRRRRLLRLLLLLLLLL.RR.LLR.LRLRLLRRRRLLLRRRRR.L.RRRRRRRRLRL.R.L.....RR.LLLLLLLL"},
}

func TestDominoPush(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := pushDominoes(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
