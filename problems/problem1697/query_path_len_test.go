package problem1697

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N              int
	Edges, Queries [][]int
	Expected       []bool
}

var Results = []Result{
	{
		3,
		[][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}},
		[][]int{{0, 1, 2}, {0, 2, 5}},
		[]bool{false, true},
	},
	{
		5,
		[][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}},
		[][]int{{0, 4, 14}, {1, 4, 13}},
		[]bool{true, false},
	},
	{
		13,
		[][]int{{9, 1, 53}, {3, 2, 66}, {12, 5, 99}, {9, 7, 26}, {1, 4, 78}, {11, 1, 62}, {3, 10, 50}, {12, 1, 71}, {12, 6, 63}, {1, 10, 63}, {9, 10, 88}, {9, 11, 59}, {1, 4, 37}, {4, 2, 63}, {0, 2, 26}, {6, 12, 98}, {9, 11, 99}, {4, 5, 40}, {2, 8, 25}, {4, 2, 35}, {8, 10, 9}, {11, 9, 25}, {10, 11, 11}, {7, 6, 89}, {2, 4, 99}, {10, 4, 63}},
		[][]int{{9, 7, 65}, {9, 6, 1}, {4, 5, 34}, {10, 8, 43}, {3, 7, 76}, {4, 2, 15}, {7, 6, 52}, {2, 0, 50}, {7, 6, 62}, {1, 0, 81}, {4, 5, 35}, {0, 11, 86}, {12, 5, 50}, {11, 2, 2}, {9, 5, 6}, {12, 0, 95}, {10, 6, 9}, {9, 4, 73}, {6, 10, 48}, {12, 0, 91}, {9, 10, 58}, {9, 8, 73}, {2, 3, 44}, {7, 11, 83}, {5, 3, 14}, {6, 2, 33}},
		[]bool{true, false, false, true, true, false, false, true, false, true, false, true, false, false, false, true, false, true, false, true, true, true, false, true, false, false},
	},
}

func TestQueryPathLength(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := distanceLimitedPathsExist(res.N, copyIntInt(res.Edges), copyIntInt(res.Queries))
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkQueryPathLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			distanceLimitedPathsExist(res.N, copyIntInt(res.Edges), copyIntInt(res.Queries))
		}
	}
}

func TestQueryPathLengthDSU(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := distanceLimitedPathsExistDSU(res.N, copyIntInt(res.Edges), copyIntInt(res.Queries))
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkQueryPathLengthDSU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			distanceLimitedPathsExistDSU(res.N, copyIntInt(res.Edges), copyIntInt(res.Queries))
		}
	}
}

func copyIntInt(a [][]int) [][]int {
	res := make([][]int, len(a))
	for i := range a {
		res[i] = make([]int, len(a[i]))
		copy(res[i], a[i])
	}
	return res
}
