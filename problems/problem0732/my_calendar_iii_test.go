package problem0732

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    [2]int
	Expected int
}

var Results = []Result{
	{[2]int{10, 20}, 1},
	{[2]int{50, 60}, 1},
	{[2]int{10, 40}, 2},
	{[2]int{5, 15}, 3},
	{[2]int{5, 10}, 3},
	{[2]int{25, 55}, 3},
}

func TestMyCalendar3(t *testing.T) {
	assert := assert.New(t)
	calendar := Constructor()

	for _, res := range Results {
		want := res.Expected
		got := calendar.Book(res.Input[0], res.Input[1])
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
