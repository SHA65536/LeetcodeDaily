package problem0352

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Vals     []int
	Expected [][][]int
}

var Results = []Result{
	{
		[]string{"SummaryRanges", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals"},
		[]int{0, 1, 0, 3, 0, 7, 0, 2, 0, 6, 0},
		[][][]int{nil, nil, {{1, 1}}, nil, {{1, 1}, {3, 3}}, nil, {{1, 1}, {3, 3}, {7, 7}}, nil, {{1, 3}, {7, 7}}, nil, {{1, 3}, {6, 7}}},
	},
}

func TestDisjointStream(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([][][]int, len(res.Expected))
		var sum SummaryRanges
		for i := range res.Input {
			switch res.Input[i] {
			case "SummaryRanges":
				sum = Constructor()
			case "addNum":
				sum.AddNum(res.Vals[i])
			case "getIntervals":
				got[i] = sum.GetIntervals()
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
