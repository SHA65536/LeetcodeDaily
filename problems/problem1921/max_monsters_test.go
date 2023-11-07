package problem1921

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Dist     []int
	Speed    []int
	Expected int
}

var TestCases = []TestCase{
	{[]int{1, 3, 4}, []int{1, 1, 1}, 3},
	{[]int{1, 1, 2, 3}, []int{1, 1, 1, 1}, 1},
	{[]int{3, 2, 4}, []int{5, 3, 2}, 1},
	{[]int{60, 72, 61, 18, 75, 42, 9, 20, 98, 52, 35, 61, 76, 80, 18, 37}, []int{7, 1, 5, 6, 8, 10, 8, 7, 4, 7, 6, 5, 8, 5, 9, 8}, 3},
}

func TestEliminateMonsters(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var want = tc.Expected
		tD, tS := make([]int, len(tc.Dist)), make([]int, len(tc.Dist))
		copy(tD, tc.Dist)
		copy(tS, tc.Speed)
		got := eliminateMaximum(tD, tS)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkEliminateMonsters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			tD, tS := make([]int, len(tc.Dist)), make([]int, len(tc.Dist))
			copy(tD, tc.Dist)
			copy(tS, tc.Speed)
			eliminateMaximum(tD, tS)
		}
	}
}

func TestEliminateMonstersNoSort(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var want = tc.Expected
		tD, tS := make([]int, len(tc.Dist)), make([]int, len(tc.Dist))
		copy(tD, tc.Dist)
		copy(tS, tc.Speed)
		got := eliminateMaximumNoSort(tD, tS)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkEliminateMonstersNoSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			tD, tS := make([]int, len(tc.Dist)), make([]int, len(tc.Dist))
			copy(tD, tc.Dist)
			copy(tS, tc.Speed)
			eliminateMaximumNoSort(tD, tS)
		}
	}
}
