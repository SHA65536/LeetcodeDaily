package problem0630

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected int
}

var Results = []Result{
	{
		[][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}},
		3,
	},
	{
		[][]int{{1, 2}},
		1,
	},
	{
		[][]int{{3, 2}, {4, 3}},
		0,
	},
	{
		[][]int{{5, 15}, {3, 19}, {6, 7}, {2, 10}, {5, 16}, {8, 14}, {10, 11}, {2, 19}},
		5,
	},
	{
		[][]int{{5, 5}, {4, 6}, {2, 6}},
		2,
	},
}

func TestScheduleCourse(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := scheduleCourse(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
