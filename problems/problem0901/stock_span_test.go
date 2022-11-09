package problem0901

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
	{
		[]int{100, 80, 60, 70, 60, 75, 85},
		[]int{1, 1, 1, 2, 1, 4, 6},
	},
	{
		[]int{85, 157, 77, 116, 191, 177, 157, 79, 159, 108, 29, 115, 66, 92, 162, 189, 91, 117, 172, 147, 15, 125, 94, 3, 15, 128, 142, 92, 24, 41, 76, 128, 113, 178, 113, 163, 86, 36, 180, 136, 144, 168, 87, 119, 179, 115, 171, 167, 78, 47, 15, 53, 138, 91, 71, 119, 54, 146, 112, 184, 124, 42, 83, 184, 130, 76, 26, 4, 47, 179, 3, 16, 119, 186, 82, 118, 199, 87, 162, 5, 90, 74, 179, 159, 112, 22, 173, 182, 96, 97, 112, 153, 59, 176, 39, 77, 106, 155, 141, 62},
		[]int{1, 2, 1, 2, 5, 1, 1, 1, 3, 1, 1, 3, 1, 2, 9, 11, 1, 2, 3, 1, 1, 2, 1, 1, 2, 6, 7, 1, 1, 2, 3, 5, 1, 18, 1, 2, 1, 1, 23, 1, 2, 3, 1, 2, 6, 1, 2, 1, 1, 1, 1, 3, 5, 1, 1, 3, 1, 10, 1, 44, 1, 1, 2, 48, 1, 1, 1, 1, 3, 6, 1, 2, 3, 58, 1, 2, 77, 1, 2, 1, 2, 1, 6, 1, 1, 1, 4, 11, 1, 2, 3, 4, 1, 6, 1, 2, 3, 4, 1, 1},
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50},
	},
}

func TestStockSpan(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]int, len(want))
		spanner := Constructor()
		for i := range res.Input {
			got[i] = spanner.Next(res.Input[i])
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkStockSpan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			spanner := Constructor()
			for j := range res.Input {
				spanner.Next(res.Input[j])
			}
		}
	}
}

func TestStockSpanOpt(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]int, len(want))
		spanner := ConstructorOpt()
		for i := range res.Input {
			got[i] = spanner.NextOpt(res.Input[i])
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkStockSpanOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			spanner := ConstructorOpt()
			for j := range res.Input {
				spanner.NextOpt(res.Input[j])
			}
		}
	}
}
