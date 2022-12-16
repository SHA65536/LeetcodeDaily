package problem0232

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Cmds    []string
	Vals    []int
	RetVal  []int
	RetBool []bool
}

var Results = []Result{
	{
		Cmds:    []string{"push", "push", "peek", "pop", "empty"},
		Vals:    []int{1, 2, 0, 0, 0},
		RetVal:  []int{0, 0, 1, 1, 0},
		RetBool: []bool{false, false, false, false, false},
	},
	{
		Cmds:    []string{"push", "push", "peek", "pop", "pop", "empty"},
		Vals:    []int{2, 1, 0, 0, 0, 0},
		RetVal:  []int{0, 0, 2, 2, 1, 0},
		RetBool: []bool{false, false, false, false, false, true},
	},
}

func TestLongestCommonSubsequence(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		q := Constructor()
		exVal := make([]int, len(res.Cmds))
		exBool := make([]bool, len(res.Cmds))
		for idx, cmd := range res.Cmds {
			switch cmd {
			case "pop":
				exVal[idx] = q.Pop()
			case "push":
				q.Push(res.Vals[idx])
			case "peek":
				exVal[idx] = q.Peek()
			case "empty":
				exBool[idx] = q.Empty()
			}
		}
		assert.Equal(res.RetVal, exVal)
		assert.Equal(res.RetBool, exBool)
	}
}
