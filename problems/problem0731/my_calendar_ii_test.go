package problem0731

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [2]int
	Expected bool
}

var Results = []Result{
	{[2]int{10, 20}, true},
	{[2]int{50, 60}, true},
	{[2]int{10, 40}, true},
	{[2]int{5, 15}, false},
	{[2]int{5, 10}, true},
	{[2]int{25, 55}, true},
}

func TestKthSmallest(t *testing.T) {
	assert := assert.New(t)
	calendar := Constructor()

	for _, res := range Results {
		want := res.Expected
		got := calendar.Book(res.Input[0], res.Input[1])
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
