package problem0081

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Target   int
	Expected bool
}

var Results = []Result{
	{[]int{4, 5, 6, 7, 0, 0, 1, 2}, 0, true},
	{[]int{4, 4, 5, 6, 7, 0, 1, 2}, 3, false},
	{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
	{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
	{[]int{1}, 0, false},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}, 19, true},
	createArray(1000),
	createArray(100000),
	createArray(10000000),
}

func createArray(length int) Result {
	var res = Result{
		Input:    make([]int, length),
		Target:   rand.Intn(length),
		Expected: true,
	}
	var pivot = rand.Intn(length)
	for i := range res.Input {
		res.Input[(i+pivot)%length] = i
	}
	return res
}

/*
func TestSearchInRotatedII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := search(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSearchInRotatedII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			search(res.Input, res.Target)
		}
	}
}
*/

func TestLinearSearchII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := linearSearch(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkLinearSearchII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			linearSearch(res.Input, res.Target)
		}
	}
}
