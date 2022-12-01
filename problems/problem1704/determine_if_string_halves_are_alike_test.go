package problem1704

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected bool
}

var Results = []Result{
	{"book", true},
	{"textbook", false},
	{"aaaa", true},
}

func TestStringHalvesAreAlike(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := halvesAreAlike(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
