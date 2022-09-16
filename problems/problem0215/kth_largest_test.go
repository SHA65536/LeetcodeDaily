package problem0215

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	K        int
	Expected int
}

var Results = []Result{
	{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
	{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	{[]int{45, 53, 42, 88, 92, 70, 83, 18, 44, 13, 56, 58, 35, 46, 58, 1, 13, 49, 74, 62, 3, 75, 80, 97, 7, 78, 92, 70, 46, 38, 69, 55, 53, 37, 55, 70, 75, 35, 56, 2, 68, 30, 88, 22, 60, 7, 99, 55, 7, 52, 68, 26, 74, 65, 100, 15, 6, 48, 92, 21, 70, 38, 72, 34, 59, 84, 35, 3, 35, 66, 78, 69, 30, 52, 78, 83, 39, 78, 29, 26, 67, 74, 95, 96, 29, 39, 67, 27, 10, 18, 70, 39, 25, 47, 91, 70, 71, 64, 24, 38}, 2, 99},
	{[]int{45, 53, 42, 88, 92, 70, 83, 18, 44, 13, 56, 58, 35, 46, 58, 1, 13, 49, 74, 62, 3, 75, 80, 97, 7, 78, 92, 70, 46, 38, 69, 55, 53, 37, 55, 70, 75, 35, 56, 2, 68, 30, 88, 22, 60, 7, 99, 55, 7, 52, 68, 26, 74, 65, 100, 15, 6, 48, 92, 21, 70, 38, 72, 34, 59, 84, 35, 3, 35, 66, 78, 69, 30, 52, 78, 83, 39, 78, 29, 26, 67, 74, 95, 96, 29, 39, 67, 27, 10, 18, 70, 39, 25, 47, 91, 70, 71, 64, 24, 38}, 15, 80},
}

func TestFindKthLargestSort(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		inputCopy := make([]int, len(res.Input))
		copy(inputCopy, res.Input)
		got := findKthLargestSort(inputCopy, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFindKthLargestSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			inputCopy := make([]int, len(res.Input))
			copy(inputCopy, res.Input)
			findKthLargestSort(inputCopy, res.K)
		}
	}
}

func TestFindKthLargest(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		inputCopy := make([]int, len(res.Input))
		copy(inputCopy, res.Input)
		got := findKthLargest(inputCopy, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			inputCopy := make([]int, len(res.Input))
			copy(inputCopy, res.Input)
			findKthLargest(inputCopy, res.K)
		}
	}
}
