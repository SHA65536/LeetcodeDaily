package problem0315

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
	{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
	{[]int{-1}, []int{0}},
	{[]int{-1, -1}, []int{0, 0}},
	{[]int{1, 2, 3, 4, 5, 6}, []int{0, 0, 0, 0, 0, 0}},
	{[]int{6, 5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1, 0}},
	{[]int{66, 18, 28, 421, -335, 215, 160, 47, 266, 434, 19, -40, -316, 434, 366, 315, -454, 172, -431, 307, -275, -179, -349, 485, -457, -64, 61, 23, 228, 33, -258, -214, -254, 376, -20, 71, -371, 80, 158, -124, -461, -274, -208, -267, -288, 267, 435, -207, 266, 150, -223, 326, -70, -17, -137, 475, -375, -216, -266, 374, -119, 401, -40, 6, -254, 379, 130, 467, -260, -123, 426, -332, -328, 41, 168, -234, -152, -499, -398, -45, 359, -270, -148, -351, 221, 234, -403, 222, -246, -170, 423, 462, -51, -426, 174, 80, 152, 8, -355, 380}, []int{60, 52, 54, 87, 13, 68, 64, 55, 70, 84, 51, 45, 15, 81, 73, 70, 3, 60, 3, 66, 14, 29, 10, 76, 2, 35, 45, 42, 55, 42, 17, 23, 17, 57, 35, 39, 6, 38, 42, 27, 1, 10, 19, 11, 9, 42, 50, 17, 40, 33, 15, 38, 22, 25, 19, 44, 4, 14, 9, 33, 17, 34, 19, 19, 10, 29, 21, 32, 9, 14, 28, 6, 6, 15, 17, 8, 9, 0, 2, 9, 16, 4, 6, 3, 10, 11, 1, 9, 2, 2, 8, 8, 2, 0, 4, 2, 2, 1, 0, 0}},
}

func TestCountSmallerNaive(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countSmallerNaive(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCountSmallerNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			countSmallerNaive(res.Input)
		}
	}
}

func TestCountSmallerInsertion(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countSmallerInsertion(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCountSmallerInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			countSmallerInsertion(res.Input)
		}
	}
}

/*func TestCountSmaller(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countSmaller(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCountSmaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			countSmaller(res.Input)
		}
	}
}*/
