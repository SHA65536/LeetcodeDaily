package problem0295

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Cmds     []string
	Input    []int
	Expected []float64
}

var Results = []Result{
	{
		[]string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
		[]int{0, 1, 2, 0, 3, 0},
		[]float64{0, 0, 0, 1.5, 0, 2.0},
	},
}

func TestStreamMedian(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		var mf MedianFinder
		want := res.Expected
		got := make([]float64, len(res.Input))
		for i := range res.Cmds {
			switch res.Cmds[i] {
			case "MedianFinder":
				mf = Constructor()
			case "addNum":
				mf.AddNum(res.Input[i])
			case "findMedian":
				got[i] = mf.FindMedian()
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkStreamMedian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			var mf MedianFinder
			for i := range res.Cmds {
				switch res.Cmds[i] {
				case "MedianFinder":
					mf = Constructor()
				case "addNum":
					mf.AddNum(res.Input[i])
				case "findMedian":
					mf.FindMedian()
				}
			}
		}
	}
}

func TestStreamMedianHeap(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		var mf MedianFinderHeap
		want := res.Expected
		got := make([]float64, len(res.Input))
		for i := range res.Cmds {
			switch res.Cmds[i] {
			case "MedianFinder":
				mf = ConstructorHeap()
			case "addNum":
				mf.AddNum(res.Input[i])
			case "findMedian":
				got[i] = mf.FindMedian()
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkStreamMedianHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			var mf MedianFinderHeap
			for i := range res.Cmds {
				switch res.Cmds[i] {
				case "MedianFinder":
					mf = ConstructorHeap()
				case "addNum":
					mf.AddNum(res.Input[i])
				case "findMedian":
					mf.FindMedian()
				}
			}
		}
	}
}
