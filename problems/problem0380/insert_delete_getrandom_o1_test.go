package problem0380

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Cmds          []string
	Vals          []int
	ExpectedVals  []int
	ExpectedBools []bool
}

var Results = []Result{
	{
		Cmds:          []string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"},
		Vals:          []int{0, 1, 2, 2, 0, 1, 2, 0},
		ExpectedVals:  []int{0, 0, 0, 0, 1, 0, 0, 1},
		ExpectedBools: []bool{false, true, false, true, false, true, false, false},
	},
}

func TestRandomizedSet(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		gotNums := make([]int, len(res.ExpectedVals))
		gotBools := make([]bool, len(res.ExpectedBools))
		var rSet RandomizedSet = Constructor()
		for i := range res.Cmds {
			switch res.Cmds[i] {
			case "insert":
				gotBools[i] = rSet.Insert(res.Vals[i])
			case "remove":
				gotBools[i] = rSet.Remove(res.Vals[i])
			case "getRandom":
				gotNums[i] = rSet.GetRandom()
				if gotNums[i] != 0 {
					gotNums[i] = 1
				}
			}
		}
		assert.Equal(gotNums, res.ExpectedVals, fmt.Sprintf("%+v", res))
		assert.Equal(gotBools, res.ExpectedBools, fmt.Sprintf("%+v", res))
	}
}
