package problem2558

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	K        int
	Expected int64
}

var Results = []Result{
	{[]int{25, 64, 9, 4, 100}, 4, 29},
	{[]int{1, 1, 1, 1}, 4, 4},
}

func TestTakeGiftsFromPiles(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := pickGifts(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
