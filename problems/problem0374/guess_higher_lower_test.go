package problem0374

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected int
}

var Results = []Result{
	{10, 6},
	{1, 1},
	{12, 4},
}

func TestGuessHigherLower(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		n := res.Expected
		want := res.Expected
		guess := func(i int) int {
			if i == n {
				return 0
			} else if i > n {
				return -1
			} else {
				return 1
			}
		}
		got := guessNumber(res.Input, guess)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
