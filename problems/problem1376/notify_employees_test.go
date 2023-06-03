package problem1376

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N, Head          int
	Managers, Infrom []int
	Expected         int
}

var TestCases = []TestCase{
	{
		N: 1, Head: 0,
		Managers: []int{-1},
		Infrom:   []int{0},
		Expected: 0,
	},

	{
		N: 6, Head: 2,
		Managers: []int{2, 2, -1, 2, 2, 2},
		Infrom:   []int{0, 0, 1, 0, 0, 0},
		Expected: 1,
	},
	{
		N: 6, Head: 0,
		Managers: []int{-1, 0, 0, 0, 3, 1},
		Infrom:   []int{2, 3, 0, 4, 0, 0},
		Expected: 6,
	},
}

func TestNotifyEmployees(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := numOfMinutes(tc.N, tc.Head, tc.Managers, tc.Infrom)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkNotifyEmployees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			numOfMinutes(tc.N, tc.Head, tc.Managers, tc.Infrom)
		}
	}
}
