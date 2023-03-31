package problem1444

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	K        int
	Expected int
}

var Results = []Result{
	{[]string{"A..", "AAA", "..."}, 3, 3},
	{[]string{"A..", "AA.", "..."}, 3, 1},
	{[]string{"A..", "A..", "..."}, 1, 1},
}

func TestCutPizza(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := ways(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
