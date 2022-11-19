package problem1629

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Keys     string
	Expected byte
}

var Results = []Result{
	{
		[]int{9, 29, 49, 50},
		"cbcd",
		'c',
	},
	{
		[]int{12, 23, 36, 46, 62},
		"spuda",
		'a',
	},
}

func TestSlowestKey(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := slowestKey(res.Input, res.Keys)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
