package problem2390

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"leet**cod*e", "lecoe"},
	{"erase*****", ""},
}

func TestRemoveStars(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := removeStars(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
