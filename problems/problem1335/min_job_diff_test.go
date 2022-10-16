package problem1335

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Days     int
	Expected int
}

var Results = []Result{
	{[]int{6, 5, 4, 3, 2, 1}, 2, 7},
	{[]int{9, 9, 9}, 4, -1},
	{[]int{1, 1, 1}, 3, 3},
	{[]int{2, 5, 6, 4, 7, 8, 96, 21, 4, 3, 4, 5, 7, 8, 66, 21, 1, 4, 5, 6, 3, 3, 5}, 5, 109},
}

func TestMinJobDifficulty(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minDifficulty(res.Input, res.Days)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMinJobDifficulty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			minDifficulty(res.Input, res.Days)
		}
	}
}

func TestMinJobDifficulty1D(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minDifficulty1D(res.Input, res.Days)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMinJobDifficulty1D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			minDifficulty1D(res.Input, res.Days)
		}
	}
}
