package problem0542

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [][]int
	Expected [][]int
}

var Results = []Result{
	{[][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	},
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}},
	{[][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	},
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 2, 1},
		}},
}

func TestUpdateMatrix(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := updateMatrix(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkUpdateMatrix(b *testing.B) {
	var mat = make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		mat[i] = make([]int, 1000)
		for j := 0; j < 1000; j++ {
			mat[i][j] = 1
		}
	}
	mat[999][999] = 0
	for i := 0; i < b.N; i++ {
		updateMatrix(mat)
	}
}

func TestUpdateMatrixOpt(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := updateMatrixOpt(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkUpdateMatrixOpt(b *testing.B) {
	var mat = make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		mat[i] = make([]int, 1000)
		for j := 0; j < 1000; j++ {
			mat[i][j] = 1
		}
	}
	mat[999][999] = 0
	for i := 0; i < b.N; i++ {
		updateMatrixOpt(mat)
	}
}
