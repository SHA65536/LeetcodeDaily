package problem1817

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	K        int
	Expected []int
}

var Results = []Result{
	{[][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5, []int{0, 2, 0, 0, 0}},
	{[][]int{{1, 1}, {2, 2}, {2, 3}}, 4, []int{1, 1, 0, 0}},
}

func TestUserActivity(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findingUsersActiveMinutes(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkUserActivity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			findingUsersActiveMinutes(res.Input, res.K)
		}
	}
}
