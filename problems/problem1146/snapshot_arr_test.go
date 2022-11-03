package problem1146

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Cmds     []string
	Inputs   [][]int
	Expected []int
}

var Results = []Result{
	{
		[]string{"SnapshotArray", "set", "snap", "set", "get"},
		[][]int{{3}, {0, 5}, {}, {0, 6}, {0, 0}},
		[]int{0, 0, 0, 0, 5},
	},
	{
		[]string{"SnapshotArray", "set", "set", "set", "snap", "get", "snap"},
		[][]int{{1}, {0, 4}, {0, 16}, {0, 13}, {}, {0, 0}, {}},
		[]int{0, 0, 0, 0, 0, 13, 1},
	},
	{
		[]string{"SnapshotArray", "snap", "snap", "get", "set", "snap", "set"},
		[][]int{{4}, {}, {}, {3, 1}, {2, 4}, {}, {1, 4}},
		[]int{0, 0, 1, 0, 0, 2, 0},
	},
	{
		[]string{"SnapshotArray", "set", "snap", "snap", "set", "set", "get", "get", "get"},
		[][]int{{3}, {1, 6}, {}, {}, {1, 19}, {0, 4}, {2, 1}, {2, 0}, {0, 1}},
		[]int{0, 0, 0, 1, 0, 0, 0, 0, 0},
	},
}

func TestSnapshotArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		var snaparr SnapshotArray
		var got = make([]int, len(res.Cmds))
		for i := range res.Cmds {
			switch res.Cmds[i] {
			case "SnapshotArray":
				snaparr = Constructor(res.Inputs[i][0])
			case "set":
				snaparr.Set(res.Inputs[i][0], res.Inputs[i][1])
			case "snap":
				got[i] = snaparr.Snap()
			case "get":
				got[i] = snaparr.Get(res.Inputs[i][0], res.Inputs[i][1])
			}
		}
		assert.Equal(res.Expected, got, fmt.Sprintf("%+v", res))
	}
}
