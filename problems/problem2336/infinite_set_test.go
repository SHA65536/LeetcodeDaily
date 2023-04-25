package problem2336

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Cmds     []string
	Values   []int
	Expected []int
}

var Results = []Result{
	{
		Cmds:     []string{"SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"},
		Values:   []int{0, 2, 0, 0, 0, 1, 0, 0, 0},
		Expected: []int{0, 0, 1, 2, 3, 0, 1, 4, 5},
	},
	{
		Cmds:     []string{"SmallestInfiniteSet", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "popSmallest"},
		Values:   []int{0, 0, 1, 0, 0, 0, 2, 3, 0, 0},
		Expected: []int{0, 1, 0, 1, 2, 3, 0, 0, 2, 3},
	},
}

func TestSmallestInInfiniteSet(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		var set SmallestInfiniteSet
		want := res.Expected
		got := make([]int, len(want))
		for i, cmd := range res.Cmds {
			switch cmd {
			case "SmallestInfiniteSet":
				set = Constructor()
			case "addBack":
				set.AddBack(res.Values[i])
			case "popSmallest":
				got[i] = set.PopSmallest()
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", got))
	}
}
