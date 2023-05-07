package problem1964

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected []int
}

var Results = []Result{
	{[]int{1, 2, 3, 2}, []int{1, 2, 3, 3}},
	{[]int{2, 2, 1}, []int{1, 2, 1}},
	{[]int{3, 1, 5, 6, 4, 2}, []int{1, 1, 2, 3, 2, 2}},
}

func TestObstacleCourses(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		temp := make([]int, len(res.Input))
		copy(temp, res.Input)
		got := longestObstacleCourseAtEachPosition(temp)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkTestObstacleCourses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			temp := make([]int, len(res.Input))
			copy(temp, res.Input)
			longestObstacleCourseAtEachPosition(temp)
		}
	}
}
