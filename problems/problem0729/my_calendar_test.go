package problem0729

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
	{[2]int{47, 50}, true},
	{[2]int{33, 41}, true},
	{[2]int{39, 45}, false},
	{[2]int{33, 42}, false},
	{[2]int{25, 32}, true},
	{[2]int{26, 35}, false},
	{[2]int{19, 25}, true},
	{[2]int{3, 8}, true},
	{[2]int{8, 13}, true},
	{[2]int{18, 27}, false},
}

func TestMyCalendar(t *testing.T) {
	assert := assert.New(t)
	calendar := Constructor()

	for _, res := range Results {
		want := res.Expected
		got := calendar.Book(res.Input[0], res.Input[1])
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
