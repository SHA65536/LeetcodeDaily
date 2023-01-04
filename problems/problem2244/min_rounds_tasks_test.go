package problem2244

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected int
}

var Results = []Result{
	{[]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}, 4},
	{[]int{2, 3, 3}, -1},
}

func TestMinRoundsForTasks(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minimumRounds(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
