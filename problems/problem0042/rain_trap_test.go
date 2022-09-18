package problem0042

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
	{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	{[]int{4, 2, 0, 3, 2, 5}, 9},
	{[]int{41, 45, 99, 17, 30, 93, 92, 21, 89, 68, 45, 15, 26, 42, 62, 88, 67, 93, 0, 56, 94, 31, 71, 5, 93, 72, 12, 59, 46, 53, 32, 58, 65, 45, 23, 78, 8, 68, 49, 66, 57, 77, 89, 89, 94, 45, 62, 67, 48, 31, 53, 68, 87, 57, 9, 67, 22, 71, 51, 15, 1, 84, 26, 27, 12, 7, 59, 79, 19, 97, 21, 41, 37, 65, 20, 0, 10, 87, 36, 25, 10, 97, 94, 50, 13, 55, 92, 82, 77, 1, 18, 57, 52, 93, 52, 33, 84, 87, 49, 100}, 4564},
}

func TestRainWaterTrap(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := trap(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
